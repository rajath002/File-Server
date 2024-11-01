package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap  map[string]string
	VideoFiles []string
}