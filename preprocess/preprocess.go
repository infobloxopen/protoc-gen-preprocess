package plugin

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	prep "github.com/infobloxopen/protoc-gen-preprocess/options"
)

func init() {
	generator.RegisterPlugin(NewPreprocessor())
}

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
		opts := getMessageOptions(message)
		if opts != nil || containsFieldPreprocessOptions(message) {
			p.generateProto3Message(message, opts)
		}
	}
}

func (p *preprocessor) generateProto3Message(message *generator.Descriptor, messageOptions *prep.PreprocessMessageOptions) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, ccTypeName, `) Preprocess() error {`)
	p.In()
	for _, field := range message.Field {
		fieldOptions := getFieldOptions(field)
		if fieldOptions == nil && !field.IsMessage() && messageOptions == nil {
			continue
		}
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "m." + fieldName
		if field.IsString() {
			p.generateStringPreprocessor(variableName, []prepOptions{messageOptions, fieldOptions}, field.IsRepeated())
		}
	}
	p.Out()
	p.P()
	p.P(`return nil`)
	p.P(`}`)
	p.P()
}

func (p *preprocessor) generateStringPreprocessor(variableName string, opts []prepOptions, repeated bool) {
	p.P()
	strMethods := make(map[string]prep.PreprocessString_Methods)
	for _, v := range opts {
		if str := v.GetString_(); str != nil {
			for _, m := range str.Methods {
				strMethods[m.String()] = m
			}
			if str.GetTrimSpace() {
				strMethods["trim_space"] = prep.PreprocessString_trim
			}
		}
	}
	if repeated {
		p.P(`for i, s := range `, variableName, `{`)
		p.In()
		for _, method := range strMethods {
			p.P(variableName, `[i] = `, p.stringsPkg.Use(), stringMethods[method], `(s)`)
		}
		p.Out()
		p.P(`}`)
	} else {
		for _, method := range strMethods {
			p.P(variableName, ` = `, p.stringsPkg.Use(), stringMethods[method], `(`, variableName, `)`)
		}
	}
}
