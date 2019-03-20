package main

import (
	"github.com/gogo/protobuf/vanity/command"
	plugin "github.com/infobloxopen/protoc-gen-preprocess/preprocess"
)

func main() {
	response := command.GeneratePlugin(command.Read(), plugin.NewPreprocessor(), ".pb.preprocess.go")
	command.Write(response)
}
