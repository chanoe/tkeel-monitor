package main

import (
	"context"
	"flag"
	"github.com/tkeel-io/tkeel-monitor/api/monitoring/v1"
	"github.com/tkeel-io/tkeel-monitor/pkg/ksclient"
	"os"
	"os/signal"
	"syscall"

	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	"github.com/tkeel-io/tkeel-monitor/pkg/server"
	"github.com/tkeel-io/tkeel-monitor/pkg/service"

	// User import.
	openapi "github.com/tkeel-io/tkeel-monitor/api/openapi/v1"
	prometheusv1 "github.com/tkeel-io/tkeel-monitor/api/prometheus/v1"
)

var (
	// Name app.
	Name string
	// HTTPAddr string.
	HTTPAddr string
	// GRPCAddr string.
	GRPCAddr string
	// PromNamespace string.
	PromNamespace string
	// PromCRDName string.
	//PromCRDName string

	// KsAddr host:port
	KsAddr string
	// TKeel NS
	TKeelNamespace string
	// TKeel cluster
	TKeelCluster string
)

func init() {
	flag.StringVar(&Name, "name", getEnvStr("APP_NAME", "tkeel-monitor"), "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", getEnvStr("HTTP_ADDR", ":31234"), "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", getEnvStr("GRPC_ADDR", ":31233"), "grpc listen address.")

	flag.StringVar(&PromNamespace, "prom_namespace", getEnvStr("PROM_NAMESPACE", "keel-system"), "prometheusv1 install namespace.")

	flag.StringVar(&KsAddr, "ks_addr", getEnvStr("KS_ADDR", ""), "ks access addr.")
	flag.StringVar(&TKeelNamespace, "tkeel_namespace", getEnvStr("TKEEL_NAMESPACE", "dapr-system"), "tkeel k8s namespace.")
	flag.StringVar(&TKeelCluster, "tkeel_cluster", getEnvStr("TKEEL_CLUSTER", "default"), "tkeel k8s cluster.")
}

func main() {
	flag.Parse()

	httpSrv := server.NewHTTPServer(HTTPAddr)
	grpcSrv := server.NewGRPCServer(GRPCAddr)
	serverList := []transport.Server{httpSrv, grpcSrv}

	app := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		serverList...,
	)

	{
		// prom service
		promSvc := service.NewPrometheusService(PromNamespace)
		prometheusv1.RegisterPrometheusHTTPServer(httpSrv.Container, promSvc)

		// ks metrics
		ksCli := ksclient.NewKApisClient(ksclient.WithBaseTokenPath(KsAddr, ""),
			ksclient.WithKSNameSpace(""),
			ksclient.WithTKeelNS(TKeelNamespace),
			ksclient.WithTKeelCluster(TKeelCluster))

		ksCli.RestyClient.OnBeforeRequest(ksCli.SecretTokenBeforeReq)
		ksSvc := service.NewKsMetricsService(ksCli)
		ksSrv := v1.NewKSMetricsServer(ksSvc)
		v1.RegisterKSMetricsHTTPServer(httpSrv.Container, ksSrv)

		OpenapiSrv := service.NewOpenapiService()
		openapi.RegisterOpenapiHTTPServer(httpSrv.Container, OpenapiSrv)
		openapi.RegisterOpenapiServer(grpcSrv.GetServe(), OpenapiSrv)
	}

	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := app.Stop(context.TODO()); err != nil {
		panic(err)
	}
}

func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}
