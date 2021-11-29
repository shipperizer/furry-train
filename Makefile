.PHONY=build proto proto-deps

GO111MODULE?=on
GOPRIVATE?=github.com/shipperizer/*
CGO_ENABLED?=0
GOOS?=linux
GOARCH?=amd64
GO?=go
GIT?=git
PROTOC?=protoc
INCLUDE?=/usr/include/protoc
PROTOBUF_GO_VERSION?=1.25.0
PROTOBUF_GRPC_GO_VERSION=1.1.0
APP_NAME?=web

.EXPORT_ALL_VARIABLES:

autoversion:
	$(GO) run autoversion.go

proto-deps:
	$(GO) get google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOBUF_GO_VERSION}
	$(GO) get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOBUF_GRPC_GO_VERSION}
	$(GO) get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	$(GO) get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	echo "pull googleapis/googleapis + grpc-ecosystem/grpc-gateway for extra protobuf files"
	$(GIT) clone git@github.com:googleapis/googleapis.git || true
	$(GIT) clone git@github.com:grpc-ecosystem/grpc-gateway.git || true

proto:
	$(PROTOC) -I=$(INCLUDE) -I=. -I=./googleapis -I=./grpc-gateway --go_opt=module=github.com/shipperizer/furry-train --go_out=. --go-grpc_opt=module=github.com/shipperizer/furry-train --go-grpc_out=. check/check.proto
	$(PROTOC) -I=$(INCLUDE) -I=. -I=./googleapis -I=./grpc-gateway --grpc-gateway_out=. --grpc-gateway_opt=module=github.com/shipperizer/furry-train --openapiv2_out=./docs --openapiv2_opt=fqn_for_openapi_name=true,allow_merge=true check/check.proto

build:
	$(MAKE) -C cmd/$(APP_NAME) build
