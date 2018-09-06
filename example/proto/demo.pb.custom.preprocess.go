package proto

// Preprocess will be called by interceptor
func (m *Custom) Preprocess() error {
	m.DoItYourself = "Changed Internals"
	return nil
}
