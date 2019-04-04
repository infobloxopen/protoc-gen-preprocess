package plugin

import (
	"sort"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	prep "github.com/infobloxopen/protoc-gen-preprocess/options"
)

type preprocessor struct {
	*generator.Generator
	generator.PluginImports
	packageName string
	stringsPkg  generator.Single
	file        *generator.FileDescriptor
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
	p.file = file
	p.stringsPkg = p.NewImport("strings")
	for _, message := range file.Messages() {
		if message.GetOptions().GetMapEntry() {
			continue
		}
		p.generateProto3Message(message, getMessageOptions(message))
	}
}

func (p *preprocessor) generateProto3Message(message *generator.Descriptor, messageOptions *prep.PreprocessMessageOptions) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, ccTypeName, `) Preprocess() error {`)
	p.In()
	for _, field := range message.Field {
		if p.IsMap(field) {
			continue
		}
		fieldOptions := getFieldOptions(field)
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "m." + fieldName
		if field.IsString() {
			p.generateStringPreprocessor(variableName, []prepOptions{messageOptions, fieldOptions}, field.IsRepeated())
		} else if field.IsMessage() && p.getPackageMessage(field.GetTypeName()) != nil {
			p.generatePreprocessCall("m."+fieldName, field.IsRepeated())
		}
	}
	p.Out()
	p.P()
	p.P(`return nil`)
	p.P(`}`)
	p.P()
}

func (p *preprocessor) getPackageMessage(t string) *descriptor.DescriptorProto {
	pkg := "." + p.file.GetPackage() + "."
	if strings.HasPrefix(t, pkg) {
		return p.file.GetMessage(strings.TrimPrefix(t, pkg))
	}

	return nil
}

func (p *preprocessor) generateStringPreprocessor(variableName string, opts []prepOptions, repeated bool) {
	p.P()
	strMethods := make(map[string]int)

	for _, v := range opts {
		if str := v.GetString_(); str != nil {
			for _, m := range str.Methods {
				switch m {
				case prep.PreprocessString_clear:
					strMethods = make(map[string]int)
				case prep.PreprocessString_none:
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
		p.P(`for i := range `, variableName, `{`)
		p.In()
		for _, method := range strOrder {
			p.P(variableName, `[i] = `, p.stringsPkg.Use(), stringMethods[method], `(`, variableName, `[i])`)
		}
		p.Out()
		p.P(`}`)
	} else {
		for _, method := range strOrder {
			p.P(variableName, ` = `, p.stringsPkg.Use(), stringMethods[method], `(`, variableName, `)`)
		}
	}
}

func (p *preprocessor) generatePreprocessCall(variableName string, repeated bool) {
	p.P()

	if repeated {
		p.P(`for _, v := range `, variableName, `{`)
		p.P(`if v != nil {`)
		p.P(`v.Preprocess()`)
		p.P(`}`)
		p.P(`}`)
	} else {
		p.P(`if `, variableName, ` != nil {`)
		p.P(variableName, `.Preprocess()`)
		p.P(`}`)
	}
}
