package proto

import (
	strings "strings"
)

func (m *Demo) Preprocess() error {

	m.PreprocessedField = strings.TrimSpace(m.PreprocessedField)
	m.PreprocessedField = strings.ToLower(m.PreprocessedField)

	for i := range m.PreprocessedRepeatedField {
		m.PreprocessedRepeatedField[i] = strings.TrimSpace(m.PreprocessedRepeatedField[i])
	}

	if m.Sub != nil {
		m.Sub.Preprocess()
	}

	for _, v := range m.SubRepeated {
		if v != nil {
			v.Preprocess()
		}
	}

	if m.Internal != nil {
		m.Internal.Preprocess()
	}

	for _, v := range m.InternalRepeated {
		if v != nil {
			v.Preprocess()
		}
	}

	for i := range m.PreprocessRepeatedField2 {
		m.PreprocessRepeatedField2[i] = strings.TrimSpace(m.PreprocessRepeatedField2[i])
		m.PreprocessRepeatedField2[i] = strings.ToLower(m.PreprocessRepeatedField2[i])
	}

	for i := range m.NonepreprocessRepeated {
		m.NonepreprocessRepeated[i] = strings.TrimSpace(m.NonepreprocessRepeated[i])
	}

	for i := range m.Clearinheritance {
		m.Clearinheritance[i] = strings.ToLower(m.Clearinheritance[i])
	}

	return nil
}
func (m *Demo_Internal) Preprocess() error {

	m.InternalString = strings.TrimSpace(m.InternalString)

	return nil
}
func (m *SubObject) Preprocess() error {

	m.StrVal = strings.TrimSpace(m.StrVal)

	return nil
}
func (m *Custom) Preprocess() error {

	return nil
}
