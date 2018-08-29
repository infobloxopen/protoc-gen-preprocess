package main

import (
	"github.com/gogo/protobuf/vanity/command"
	plugin "github.com/infobloxopen/protoc-gen-preprocess/preprocess"
)

func main() {
	req := command.Read()
	p := plugin.NewPreprocessor()
	resp := command.GeneratePlugin(req, p, ".pb.preprocess.go")
	command.Write(resp)
}
