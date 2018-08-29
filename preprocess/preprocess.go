package plugin

import (
	prep "github.com/atorgayev/protoc-gen-preprocess/options"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type preprocessor struct {
	*generator.Generator
	generator.PluginImports
	stringsPkg generator.Single
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
		p.generateProto3Message(file, message)
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

func (p *preprocessor) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
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
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, fieldOptions)
		}
	}
	p.Out()
	p.P(`return nil`)
	p.P(`}`)
	p.P()
}

func (p *preprocessor) generateStringValidator(variableName string, ccTypeName string, fieldName string, fv *prep.PreprocessFieldOptions) {
	if fv.String_ != nil {
		if fv.String_.GetTrimSpace() {
			p.P(variableName, `= `, p.stringsPkg.Use(), `.TrimSpace(`, variableName, `)`)
		}
	}
}
