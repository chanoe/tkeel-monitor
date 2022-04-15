package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	"github.com/tkeel-io/tkeel-monitor/pkg/server"
	"github.com/tkeel-io/tkeel-monitor/pkg/service"

	// User import.
	prometheus "github.com/tkeel-io/tkeel-monitor/api/prometheus/v1"

	openapi "github.com/tkeel-io/tkeel-monitor/api/openapi/v1"
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
	PromCRDName string
)

func init() {
	flag.StringVar(&Name, "name", getEnvStr("APP_NAME", "tkeel-monitor"), "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", getEnvStr("HTTP_ADDR", ":31234"), "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", getEnvStr("GRPC_ADDR", ":31233"), "grpc listen address.")

	flag.StringVar(&PromNamespace, "prom_namespace", getEnvStr("PROM_NAMESPACE", "tkeel-system"), "prometheus install namespace.")
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

	{ // User service
		ProSrv := service.NewPrometheusService(PromNamespace)
		prometheus.RegisterPrometheusHTTPServer(httpSrv.Container, ProSrv)
		prometheus.RegisterPrometheusServer(grpcSrv.GetServe(), ProSrv)

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

func getEnvBool(env string, defaultValue bool) bool {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	ret, err := strconv.ParseBool(v)
	if err != nil {
		panic(errors.Wrapf(err, "get env(%s) bool", env))
	}
	return ret
}

func getEnvInt(env string, defaultValue int) int {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	ret, err := strconv.Atoi(v)
	if err != nil {
		panic(errors.Wrapf(err, "get env(%s) int", env))
	}
	return ret
}
