package plugin

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	prep "github.com/infobloxopen/protoc-gen-preprocess/options"
)

type preprocessor struct {
	*generator.Generator
	generator.PluginImports
	packageName string
	stringsPkg  generator.Single
}

func NewPreprocessor() *preprocessor {
	p := &preprocessor{}
	return p
}

func (p *preprocessor) Name() string {
	return "preprocessor"
}

func (p *preprocessor) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *preprocessor) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.stringsPkg = p.NewImport("strings")
	for _, message := range file.Messages() {
		if containsPreprocessOptions(message) {
			p.generateProto3Message(message)
		}
	}
}

func init() {
	generator.RegisterPlugin(NewPreprocessor())
}

func getFieldOptions(field *descriptor.FieldDescriptorProto) *prep.PreprocessFieldOptions {
	if field.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(field.Options, prep.E_Field)
	if err != nil {
		return nil
	}
	opts, ok := v.(*prep.PreprocessFieldOptions)
	if !ok {
		return nil
	}
	return opts
}

func (p *preprocessor) generateProto3Message(message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, ccTypeName, `) Preprocess() error {`)
	p.In()
	for _, field := range message.Field {
		fieldOptions := getFieldOptions(field)
		if fieldOptions == nil && !field.IsMessage() {
			continue
		}
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "m." + fieldName
		repeated := field.IsRepeated()
		if field.IsString() {
			p.generateStringPreprocessor(variableName, fieldOptions, repeated)
		}
	}
	p.Out()
	p.P()
	p.P(`return nil`)
	p.P(`}`)
	p.P()
}

var stringMethods = map[prep.PreprocessString_Methods]string{
	prep.PreprocessString_trim_space: ".TrimSpace",
	prep.PreprocessString_upper:      ".ToUpper",
	prep.PreprocessString_lower:      ".ToLower",
}

func (p *preprocessor) generateStringPreprocessor(variableName string, fv *prep.PreprocessFieldOptions, repeated bool) {
	str := fv.GetString_()
	p.P()
	if repeated {
		p.P(`for i, s := range `, variableName, `{`)
		p.In()
		for _, method := range str.Methods {
			p.P(variableName, `[i] = `, p.stringsPkg.Use(), stringMethods[method], `(s)`)
		}
		p.Out()
		p.P(`}`)
	} else {
		for _, method := range str.Methods {
			p.P(variableName, ` = `, p.stringsPkg.Use(), stringMethods[method], `(`, variableName, `)`)
		}
	}
}

func containsPreprocessOptions(message *generator.Descriptor) bool {
	for _, field := range message.Field {
		fieldOptions := getFieldOptions(field)
		if fieldOptions != nil {
			return true
		}
	}
	return false
}
