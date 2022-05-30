GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

# docker tag
DOCKER_TAG ?= latest

.PHONY: init
# init env
init:
	go get -d -u  github.com/tkeel-io/tkeel-interface/openapi
	go get -d -u  github.com/tkeel-io/kit
	go get -d -u  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

	go install github.com/ocavue/protoc-gen-typescript@latest
	go install  github.com/tkeel-io/tkeel-interface/tool/cmd/artisan@latest
	go install  google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install  google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-http@latest
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-errors@latest
	go install  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --go-errors_out=paths=source_relative:. \
 	       --openapiv2_out=./api/ \
		   --openapiv2_opt=allow_merge=true \
 	       --openapiv2_opt=logtostderr=true \
 	       --openapiv2_opt=json_names_for_fields=false \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X github.com/tkeel-io/tkeel/pkg/version.GitCommit=${GIT_COMMIT} -X github.com/tkeel-io/tkeel/pkg/version.GitVersion=${GIT_VERSION} -X github.com/tkeel-io/tkeel/pkg/version.BuildDate=${BUILD_DATE} -X github.com/tkeel-io/tkeel/pkg/version.Version=${TKEEL_VERSION} -s -w" -o ./bin/ ./cmd/monitor/

.PHONY: generate
# generate
generate:
	go generate ./...


.PHONY: all
# generate all
all:
	make api;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

################################################################################
# Target: build docker image                                                   #
################################################################################
.PHONY: build-dev-container
build-dev-container:
	docker build -t tkeelio/tkeel-monitor:$(DOCKER_TAG) .
	docker push  tkeelio/tkeel-monitor:$(DOCKER_TAG)
ifeq ($(DOCKER_TAG),latest)
	docker tag tkeelio/tkeel-monitor:$(DOCKER_TAG) tkeelio/tkeel-monitor:dev
	docker push  tkeelio/tkeel-monitor:dev
endif