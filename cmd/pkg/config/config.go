package config

import "html/template"


// Holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}