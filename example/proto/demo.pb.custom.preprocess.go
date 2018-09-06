package proto

// CustomPreprocess will be called instead of generated methods
func (m *Custom) CustomPreprocess() error {
	m.DoItYourself = "Changed Internals"
	return nil
}
