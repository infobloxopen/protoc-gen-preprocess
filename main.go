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

			for _, message := range protoFile.Messages {
				if message.Desc.IsMapEntry() {
					continue
				}
				generateProto3Message(g, message)
			}
		}
	}

	return plugin.Response()
}

func generateProto3Message(g *protogen.GeneratedFile, message *protogen.Message) {
	typeName := camelCase(string(message.Desc.Name()))
	fmt.Fprintf(os.Stderr, "typeName: %s\n", typeName)

	g.P(`func (m *`, typeName, `) Preprocess() error {`)
	g.P()
	g.P("return nil")
	g.P("}")
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

func camelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
