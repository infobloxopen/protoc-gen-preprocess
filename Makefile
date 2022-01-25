GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

PROJECT_ROOT := github.com/infobloxopen/protoc-gen-preprocess

DOCKERFILE_PATH := $(CURDIR)/docker
IMAGE_REGISTRY ?= infoblox
IMAGE_VERSION  ?= dev-preprocess

# configuration for the protobuf gentool
SRCROOT_ON_HOST      := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER := /go/src/$(PROJECT_ROOT)
DOCKERPATH           := /go/src
DOCKER_RUNNER        := docker run --rm
DOCKER_RUNNER        += -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR     := infoblox/atlas-gentool:$(IMAGE_VERSION)
GENERATOR            := $(DOCKER_RUNNER) $(DOCKER_GENERATOR)

GENPREPROCESS_IMAGE      := $(IMAGE_REGISTRY)/atlas-gentool
GENPREPROCESS_DOCKERFILE := $(DOCKERFILE_PATH)/Dockerfile

default: vendor options install

.PHONY: vendor
vendor:
	go mod vendor

install:
	go install

.PHONY: gentool
gentool: vendor
	docker build -f $(GENPREPROCESS_DOCKERFILE) -t $(GENPREPROCESS_IMAGE):$(IMAGE_VERSION) .
	docker image prune -f --filter label=stage=server-intermediate

gentool-options:
	$(GENERATOR) \
		--go_out="Mgoogle/protobuf/descriptor.proto:$(DOCKERPATH)" \
		$(PROJECT_ROOT)/options/preprocess.proto

# examples related build targets

gentool-examples: vendor gentool-examples-proto gentool-examples-build

gentool-examples-proto: gentool
	$(GENERATOR) \
		-I/go/src/github.com/infobloxopen/protoc-gen-preprocess \
		-I$(SRCROOT_IN_CONTAINER) \
		--go_out="paths=source_relative:." \
		--go-grpc_out="$(DOCKERPATH)" \
		--grpc-gateway_out="logtostderr=true:$(DOCKERPATH)" \
		--preprocess_out="$(DOCKERPATH)" \
			github.com/infobloxopen/protoc-gen-preprocess/example/proto/demo.proto

gentool-examples-build: vendor
	mkdir -p .build/bin/
	docker run --rm \
		-v $(SRCROOT_ON_HOST):/go/src/$(PROJECT_ROOT) \
		golang:1.17.6-alpine \
			sh -c "cd /go/src/$(PROJECT_ROOT) && go build -o .build/bin/ $(PROJECT_ROOT)/example/"

