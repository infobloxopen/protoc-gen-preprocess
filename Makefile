
IMAGE_REGISTRY ?= infoblox
DOCKERFILE_PATH := $(CURDIR)/docker
IMAGE_VERSION  ?= dev-preprocess
GENGORM_IMAGE      := $(IMAGE_REGISTRY)/atlas-gentool
GENGORM_DOCKERFILE := $(DOCKERFILE_PATH)/Dockerfile

.PHONY: options install demo gentool

default: options install demo 

install: options
	go install

options:
	protoc -I /usr/local/include/ -I. --gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:." options/preprocess.proto  

demo: options install
	protoc -I$(GOPATH)/src -Iexample -I/usr/local/include -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--preprocess_out=./example/proto/ \
	--go_out=plugins=grpc:./example/proto/ \
	--grpc-gateway_out=./example/proto/ \
	demo.proto

gentool:
	@docker build -f $(GENGORM_DOCKERFILE) -t $(GENGORM_IMAGE):$(IMAGE_VERSION) .
	@docker tag $(GENGORM_IMAGE):$(IMAGE_VERSION) $(GENGORM_IMAGE):latest
	@docker image prune -f --filter label=stage=server-intermediate

gentool-example: gentool
	docker run --rm infoblox/atlas-gentool:dev-preprocess \
		-preprocess_out=./example/proto/ \
		--go_out=plugins=grpc:./example/proto/ \
		--grpc-gateway_out=./example/proto/ \
		demo.proto