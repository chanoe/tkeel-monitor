package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	go_restful "github.com/emicklei/go-restful"
	"github.com/tkeel-io/kit/errors"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/result"
	"github.com/tkeel-io/tkeel-monitor/pkg/service"
)

func RegisterKSMetricsHTTPServer(container *go_restful.Container, ksSvc service.KsMetricsService) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	srv := NewKSMetricsServer(ksSvc)
	ws.Route(ws.GET("/monitoring/plugins").
		To(srv.MetricsPlugins))
	ws.Route(ws.GET("/monitoring/plugins/pods").
		To(srv.PluginPods))
	ws.Route(ws.GET("/monitoring/pods/metrics").
		To(srv.PodsCpuMem))
}

type KSMetricsServer struct {
	ksSvc service.KsMetricsService
}

func NewKSMetricsServer(ksSvc service.KsMetricsService) *KSMetricsServer {
	return &KSMetricsServer{ksSvc: ksSvc}
}
func (s *KSMetricsServer) MetricsPlugins(req *go_restful.Request, resp *go_restful.Response) {
	res, err := s.ksSvc.MetricsPlugins(req.QueryParameter("name"),
		req.QueryParameter("status"),
		req.QueryParameter("order_by"),
		req.QueryParameter("page_size"),
		req.QueryParameter("page_num"))
	if err != nil {
		log.Error(err)
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}

	resData := make(map[string]interface{})
	err = json.Unmarshal(res, &resData)
	if err != nil {
		log.Error(err)
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK,
		result.Set(errors.Success.Reason, "", resData), "application/json")
	return
}

func (s *KSMetricsServer) PluginPods(req *go_restful.Request, resp *go_restful.Response) {
	plugin := req.QueryParameter("plugin")
	res, err := s.ksSvc.PluginPods(plugin)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}
	resData := make(map[string]interface{})
	err = json.Unmarshal(res, &resData)
	if err != nil {
		log.Error(err)
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK,
		result.Set(errors.Success.Reason, "", resData), "application/json")
	return
}

func (s *KSMetricsServer) PodsCpuMem(req *go_restful.Request, resp *go_restful.Response) {
	var (
		start = req.QueryParameter("start")
		end   = req.QueryParameter("end")
		step  = req.QueryParameter("step")
		times = req.QueryParameter("times")
	)
	now := time.Now().Unix()
	thirtyMAgo := now - 1800
	if start == "" {
		start = strconv.FormatInt(thirtyMAgo, 10)
		end = strconv.FormatInt(now, 10)
		step = "60s"
		times = "30"
	}
	res, err := s.ksSvc.PodsCpuMem(req.QueryParameter("plugin"), req.QueryParameter("resources"), start, end, step, times)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}
	resData := make(map[string]interface{})
	err = json.Unmarshal(res, &resData)
	if err != nil {
		log.Error(err)
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, res), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK,
		result.Set(errors.Success.Reason, "", resData), "application/json")
	return
}
