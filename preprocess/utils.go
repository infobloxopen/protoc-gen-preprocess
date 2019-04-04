package plugin

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	prep "github.com/infobloxopen/protoc-gen-preprocess/options"
)

var stringMethods = map[int]string{
	int(prep.PreprocessString_none):  "",
	int(prep.PreprocessString_trim):  ".TrimSpace",
	int(prep.PreprocessString_upper): ".ToUpper",
	int(prep.PreprocessString_lower): ".ToLower",
	int(prep.PreprocessString_clear): "",
}

type prepOptions interface {
	GetString_() *prep.PreprocessString
}

func containsFieldPreprocessOptions(message *generator.Descriptor) bool {
	for _, field := range message.Field {
		fieldOptions := getFieldOptions(field)
		if fieldOptions != nil {
			return true
		}
	}
	return false
}

func getMessageOptions(message *generator.Descriptor) *prep.PreprocessMessageOptions {
	if message.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(message.Options, prep.E_Each)
	if err != nil {
		return nil
	}
	opts, ok := v.(*prep.PreprocessMessageOptions)
	if !ok {
		return nil
	}
	return opts
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
