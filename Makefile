
CUR_DIR 			:= $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
REPO 				:= github.com/infobloxopen/protoc-gen-preprocess
SRCROOT 			:= /go/src/$(REPO)
IMAGE_REGISTRY 		?= infoblox
DOCKERFILE_PATH 	:= $(CURDIR)/docker
DOCKERPATH          := /go/src
IMAGE_VERSION  		?= dev-preprocess
GENPREP_IMAGE      	:= $(IMAGE_REGISTRY)/atlas-gentool
GENPREP_DOCKERFILE 	:= $(DOCKERFILE_PATH)/Dockerfile

.PHONY: options install demo

default: options install demo 

install: options
	go install

options:
	protoc -I /usr/local/include/ -I. --gogo_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:." options/preprocess.proto  

demo: options install
	protoc -Iexample/proto -I$(GOPATH)/src -I/usr/local/include -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--preprocess_out=./example/proto \
	--go_out=plugins=grpc:./example/proto \
	--grpc-gateway_out=./example/proto \
	demo.proto

gentool:
	@docker build -f $(GENPREP_DOCKERFILE) -t $(GENPREP_IMAGE):$(IMAGE_VERSION) .
	@docker image prune -f --filter label=stage=server-intermediate

gentool-example:
	docker run --rm -v $(CUR_DIR):$(SRCROOT) $(GENPREP_IMAGE):$(IMAGE_VERSION) \
		--preprocess_out=$(SRCROOT) \
		--go_out=plugins=grpc:$(SRCROOT) \
		--grpc-gateway_out=$(SRCROOT) \
		example/proto/demo.proto
		