default: install options demo

install:
	go install

.PHONY: options
options:
	protoc -I /usr/local/include/ -I. --go_out="Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:." options/preprocess.proto  

.PHONY: demo
demo:
	protoc -I/home/aidyn/go/src -Iexample -I/usr/local/include --preprocess_out=./example/proto/ --go_out=plugins=grpc:./example/proto/ demo.proto