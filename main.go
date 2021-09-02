package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

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

				// capture internal messages
				for _, msg := range message.Messages {
					processMessage(g, msg)
				}
			}
		}
	}

	return plugin.Response()
}

func processMessage(g *protogen.GeneratedFile, message *protogen.Message) {
	var found bool
	messageOptions := getMessageOptions(message)
	if messageOptions != nil {
		found = true
	}

	for _, field := range message.Fields {
		options := getFieldOptions(field)
		if options != nil {
			found = true
			break
		}
	}

	if found {
		generateProto3Message(g, message)
	}

	for _, internalMessage := range message.Messages {
		processMessage(g, internalMessage)
	}
}

func generateProto3Message(g *protogen.GeneratedFile, message *protogen.Message) {
	typeName := message.GoIdent.GoName
	g.P(`func (m *`, typeName, `) Preprocess() error {`)

	for _, field := range message.Fields {
		if field.Desc.IsMap() {
			continue
		}
		fieldOpts := getFieldOptions(field)
		fieldName := string(field.GoName)
		varName := "m." + fieldName
		if field.Desc.Kind().String() == "string" {
			generateStringPreprocessor(g, varName, []prepOptions{fieldOpts, getMessageOptions(message)}, field.Desc.IsList())
		} else if field.Desc.Message() != nil { // TODO: check for same package?
			generatePreprocessCall(g, varName, field.Desc.IsList())
		}
	}

	g.P()
	g.P("return nil")
	g.P("}")
}

func generatePreprocessCall(g *protogen.GeneratedFile, varName string, repeated bool) {
	g.P()

	if repeated {
		g.P(`for _, v := range `, varName, `{`)
		g.P(`if v != nil {`)
		g.P(`v.Preprocess()`)
		g.P(`}`)
		g.P(`}`)
	} else {
		g.P(`if `, varName, ` != nil {`)
		g.P(varName, `.Preprocess()`)
		g.P(`}`)
	}
}

func generateStringPreprocessor(g *protogen.GeneratedFile, varName string, opts []prepOptions, repeated bool) {
	g.P()
	strMethods := make(map[string]int)

	for _, v := range opts {
		if str := v.GetString_(); str != nil {
			for _, m := range str.Methods {
				switch m {
				case preprocess.PreprocessString_clear:
					strMethods = make(map[string]int)
				case preprocess.PreprocessString_none:
					continue
				default:
					strMethods[m.String()] = int(m)
				}
			}
		}
	}
	if len(strMethods) == 0 {
		return
	}

	strOrder := make([]int, len(strMethods))
	i := 0
	for _, v := range strMethods {
		strOrder[i] = v
		i++
	}

	sort.IntSlice(strOrder).Sort()

	if repeated {
		g.P(`for i := range `, varName, `{`)
		for _, method := range strOrder {
			g.P(varName, `[i] = `, generateImport(stringMethods[method], "strings", g), `(`, varName, `[i])`)
		}
		g.P(`}`)
	} else {
		for _, method := range strOrder {
			g.P(varName, ` = `, generateImport(stringMethods[method], "strings", g), `(`, varName, `)`)
		}
	}
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
	if v == nil {
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

var stringMethods = map[int]string{
	int(preprocess.PreprocessString_none):  "",
	int(preprocess.PreprocessString_trim):  "TrimSpace",
	int(preprocess.PreprocessString_upper): "ToUpper",
	int(preprocess.PreprocessString_lower): "ToLower",
	int(preprocess.PreprocessString_clear): "",
}

type prepOptions interface {
	GetString_() *preprocess.PreprocessString
}

func containsFieldPreprocessOptions(message *protogen.Message) bool {
	for _, field := range message.Fields {
		fieldOptions := getFieldOptions(field)
		if fieldOptions != nil {
			return true
		}
	}

	return false
}

func generateImport(name string, importPath string, g *protogen.GeneratedFile) string {
	return g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       name,
		GoImportPath: protogen.GoImportPath(importPath),
	})
}
