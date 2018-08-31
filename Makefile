.PHONY: options install demo

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