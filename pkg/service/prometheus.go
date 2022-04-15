package service

import (
	"context"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/tkeel-io/kit/log"
	pb "github.com/tkeel-io/tkeel-monitor/api/prometheus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type PrometheusService struct {
	mc   *v1.MonitoringV1Client
	pAPI promv1.API
	pb.UnimplementedPrometheusServer
}

func NewPrometheusService(namespace string) *PrometheusService {
	conf, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("in cluster config: %s", err)
	}
	mc, err := v1.NewForConfig(conf)
	if err != nil {
		log.Fatalf("new monitor client: %s", err)
	}
	l, err := mc.Prometheuses(namespace).List(context.TODO(), metav1.ListOptions{Limit: 1})
	if err != nil {
		log.Errorf("list prometheus: %s", err)
		return nil
	}
	if len(l.Items) == 0 {
		log.Fatalf("namespace(%s) not has prometheus", namespace)
	}
	url := ""
	p := l.Items[0]
	if p.Spec.ExternalURL != "" {
		url = p.Spec.ExternalURL
	} else {
		url = "http://" + p.Spec.ServiceAccountName + "." + p.ObjectMeta.Namespace + ":9090"
	}
	pc, err := api.NewClient(api.Config{
		Address: url,
	})
	if err != nil {
		log.Fatalf("new prometheus client err: %s", err)
	}
	pAPI := promv1.NewAPI(pc)
	pAPI.Buildinfo(context.TODO())
	res, err := pAPI.Buildinfo(context.TODO())
	if err != nil {
		log.Errorf("new prometheus api build.info err: %s", err)
	}
	log.Debugf("prometheus build.info:\n %v", res)
	return &PrometheusService{
		pAPI: pAPI,
		mc:   mc,
	}
}

func (s *PrometheusService) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	var (
		value model.Value
		warn  promv1.Warnings
		err   error
	)
	st := time.Unix(req.GetSt(), 0)
	et := time.Unix(req.GetEt(), 0)
	if req.GetStep() == "" {
		value, warn, err = s.pAPI.Query(ctx, req.GetQuery(), st)
	} else {
		step, err1 := time.ParseDuration(req.GetStep())
		if err1 != nil {
			log.Errorf("time step parse err: %s", err)
			return nil, pb.ResourceErrUnknown()
		}
		value, warn, err = s.pAPI.QueryRange(ctx, req.GetQuery(), promv1.Range{
			Start: st,
			End:   et,
			Step:  step,
		})
	}
	if warn != nil {
		log.Warnf("query %s warn: %v", req, warn)
	}
	return &pb.QueryResponse{
		Result: value.String(),
	}, nil
}
