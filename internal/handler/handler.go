package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// ext extracts the file extension from a given file path
func ext(filePath string) string {
	return filepath.Ext(filePath)
}

var templates *template.Template

// InitTemplates initializes the template parsing with a function map
func InitTemplates() {
	funcMap := template.FuncMap{
		"ext": ext,
	}

	var err error
	templates, err = template.New("").Funcs(funcMap).ParseGlob(filepath.Join("web/templates", "*.html"))
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
}

// RenderTemplate renders the specified template with the provided data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		log.Println("error executing template:", err)
		//ErrorPageHandler(w, "Please try again later", nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

