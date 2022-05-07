module github.com/tkeel-io/tkeel-monitor

go 1.16

require (
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/go-resty/resty/v2 v2.7.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.55.1
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/common v0.32.1
	github.com/spf13/cobra v1.4.0
	github.com/tkeel-io/kit v0.0.0-20220311032953-c8bf1e0f86cb
	github.com/tkeel-io/tkeel v0.4.1
	github.com/tkeel-io/tkeel-interface/openapi v0.0.0-20220311032517-6775df80b836
	google.golang.org/genproto v0.0.0-20220211171837-173942840c17
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
)
