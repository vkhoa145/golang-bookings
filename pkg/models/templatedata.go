package models

// TemplateData hold data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // In case of not knowing the type, we can use interface type
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
