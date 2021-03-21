.PHONY=build proto proto-deps

GO111MODULE?=on
GOPRIVATE?=github.com/shipperizer/*
CGO_ENABLED?=0
GOOS?=linux
GO?=go
PROTOC?=protoc
INCLUDE?=/usr/include/protoc
PROTOBUF_GO_VERSION?=1.25.0
PROTOBUF_GRPC_GO_VERSION=1.1.0
APP_NAME?=kafka

.EXPORT_ALL_VARIABLES:

autoversion:
	$(GO) run autoversion.go

proto-deps:
	$(GO) get google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOBUF_GO_VERSION}
	$(GO) get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOBUF_GRPC_GO_VERSION}

proto:
	$(PROTOC) -I=$(INCLUDE) -I=. --go_opt=module=github.com/shipperizer/furry-train --go_out=. types/messages.proto

build:
	$(MAKE) -C cmd/$(APP_NAME) build
