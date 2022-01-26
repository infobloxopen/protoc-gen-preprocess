package proto

import (
	strings "strings"
)

func (m *Demo) Preprocess() error {
	if err := m.testPreFunction(); err != nil {
		return err
	}

	m.PreprocessedField = strings.TrimSpace(m.PreprocessedField)
	m.PreprocessedField = strings.ToLower(m.PreprocessedField)

	for i := range m.PreprocessedRepeatedField {
		m.PreprocessedRepeatedField[i] = strings.TrimSpace(m.PreprocessedRepeatedField[i])
	}

	if m.Sub != nil {
		if err := m.Sub.Preprocess(); err != nil {
			return err
		}
	}

	for _, v := range m.SubRepeated {
		if v != nil {
			if err := v.Preprocess(); err != nil {
				return err
			}
		}
	}

	if m.Internal != nil {
		if err := m.Internal.Preprocess(); err != nil {
			return err
		}
	}

	for _, v := range m.InternalRepeated {
		if v != nil {
			if err := v.Preprocess(); err != nil {
				return err
			}
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
	if err := m.testPostFunction(); err != nil {
		return err
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
