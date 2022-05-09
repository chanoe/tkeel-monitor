package v1

import (
	"net/http"

	go_restful "github.com/emicklei/go-restful"
	"github.com/tkeel-io/kit/errors"
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

}
func (s *KSMetricsServer) PluginPods(req *go_restful.Request, resp *go_restful.Response) {
	plugin := req.QueryParameter("plugin")
	if plugin != "" {

	}
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
	resp.WriteHeaderAndJson(http.StatusOK,
		result.Set(errors.Success.Reason, "", res), "application/json")
	return
}
func (s *KSMetricsServer) PodsCpuMem(req *go_restful.Request, resp *go_restful.Response) {

}

type Metrics struct {
	Results     []interface{} `json:"results" description:"actual array of results"`
	CurrentPage int           `json:"page,omitempty" description:"current page returned"`
	TotalPages  int           `json:"total_page,omitempty" description:"total number of pages"`
	TotalItems  int           `json:"total_item,omitempty" description:"page size"`
}
type ListResult struct {
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"totalItems"`
}
