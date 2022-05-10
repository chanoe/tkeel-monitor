// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	result "github.com/tkeel-io/kit/result"
	protojson "google.golang.org/protobuf/encoding/protojson"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

var (
	_ = protojson.MarshalOptions{}
	_ = anypb.Any{}
	_ = emptypb.Empty{}
)

type PrometheusHTTPServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

type PrometheusHTTPHandler struct {
	srv PrometheusHTTPServer
}

func newPrometheusHTTPHandler(s PrometheusHTTPServer) *PrometheusHTTPHandler {
	return &PrometheusHTTPHandler{srv: s}
}

func (h *PrometheusHTTPHandler) Query(req *go_restful.Request, resp *go_restful.Response) {
	in := QueryRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.Query(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: errors.Success.Reason,
		Msg:  "",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}
	resp.AddHeader(go_restful.HEADER_ContentType, "application/json")

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func RegisterPrometheusHTTPServer(container *go_restful.Container, srv PrometheusHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newPrometheusHTTPHandler(srv)
	ws.Route(ws.GET("/prometheus").
		To(handler.Query))
}
