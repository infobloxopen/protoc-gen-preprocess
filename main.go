package main

import (
	plugin "github.com/atorgayev/protoc-gen-preprocess/preprocess"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	p := plugin.NewPreprocessor()
	resp := command.GeneratePlugin(req, p, ".pb.preprocess.go")
	command.Write(resp)
}
