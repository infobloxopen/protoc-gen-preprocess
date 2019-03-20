GOPATH ?= $(HOME)/go
SRCPATH := $(patsubst %/,%,$(GOPATH))/src

PROJECT_ROOT := github.com/infobloxopen/protoc-gen-preprocess

DOCKERFILE_PATH := $(CURDIR)/docker
IMAGE_REGISTRY ?= infoblox
IMAGE_VERSION  ?= dev-preprocess

# configuration for the protobuf gentool
SRCROOT_ON_HOST      := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER := /home/go/src/$(PROJECT_ROOT)
DOCKERPATH           := /home/go/src
DOCKER_RUNNER        := docker run --rm
DOCKER_RUNNER        += -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR     := infoblox/atlas-gentool:$(IMAGE_VERSION)
GENERATOR            := $(DOCKER_RUNNER) $(DOCKER_GENERATOR)

GENPREPROCESS_IMAGE      := $(IMAGE_REGISTRY)/atlas-gentool
GENPREPROCESS_DOCKERFILE := $(DOCKERFILE_PATH)/Dockerfile

default: vendor options install

.PHONY: vendor
vendor:
	dep ensure -vendor-only

.PHONY: vendor-update
vendor-update:
	dep ensure

install:
	go install

.PHONY: gentool
gentool:
	docker build -f $(GENPREPROCESS_DOCKERFILE) -t $(GENPREPROCESS_IMAGE):$(IMAGE_VERSION) .
	docker image prune -f --filter label=stage=server-intermediate

gentool-examples: gentool
	$(GENERATOR) \
		-I/go/src/github.com/infobloxopen/protoc-gen-preprocess \
		--go_out="plugins=grpc:$(DOCKERPATH)" \
		--grpc-gateway_out="logtostderr=true:$(DOCKERPATH)" \
		--preprocess_out="$(DOCKERPATH)" \
			example/proto/demo.proto

gentool-options:
	$(GENERATOR) \
		--gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:$(DOCKERPATH)" \
		$(PROJECT_ROOT)/options/preprocess.proto
