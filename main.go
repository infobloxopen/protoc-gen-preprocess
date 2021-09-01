package main

import (
	"fmt"
	"io/ioutil"
	"os"

	preprocess "github.com/infobloxopen/protoc-gen-preprocess/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var request pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(input, &request)
	if err != nil {
		panic(err)
	}

	opts := protogen.Options{}

	plugin, err := opts.New(&request)
	if err != nil {
		panic(err)
	}

	resp := generate(plugin)
	out, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(out))
}

func generate(plugin *protogen.Plugin) *pluginpb.CodeGeneratorResponse {
	for _, protoFile := range plugin.Files {
		var found bool
		for _, message := range protoFile.Messages {
			messageOptions := getMessageOptions(message)
			if messageOptions != nil {
				found = true
				break
			}

			for _, field := range message.Fields {
				options := getFieldOptions(field)
				if options != nil {
					found = true
					break
				}
			}
		}

		if found {
			fileName := protoFile.GeneratedFilenamePrefix + ".pb.preprocess.go"
			g := plugin.NewGeneratedFile(fileName, ".")
			g.P("package ", protoFile.GoPackageName)
		}
	}

	return plugin.Response()
}

func getMessageOptions(message *protogen.Message) *preprocess.PreprocessMessageOptions {
	options := message.Desc.Options()
	if options == nil {
		return nil
	}
	v := proto.GetExtension(options, preprocess.E_Each)
	if v == nil {
		return nil
	}

	opts, ok := v.(*preprocess.PreprocessMessageOptions)
	if !ok {
		return nil
	}

	return opts
}

func getFieldOptions(field *protogen.Field) *preprocess.PreprocessFieldOptions {
	options := field.Desc.Options()
	if options == nil {
		return nil
	}

	v := proto.GetExtension(options, preprocess.E_Field)
	if v != nil {
		return nil
	}

	opts, ok := v.(*preprocess.PreprocessFieldOptions)
	if !ok {
		return nil
	}
	return opts
}
