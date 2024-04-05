package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	if tmpl == nil {
		tmpl = template.Must(template.ParseGlob("views/*.html"))
	}
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "Golang Web Server",
		"Heading": "Hello World",
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
