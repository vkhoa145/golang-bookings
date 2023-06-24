package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Reason: This configuration file might be access from anywhere in application
// ensure this config files import 'adequate things'
// avoid: package import to each other regardless its logic
// danger: import cycle -> app will not compile

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
