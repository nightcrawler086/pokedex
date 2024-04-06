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

var i int

func GetIndex(w http.ResponseWriter, r *http.Request) {
	i++
	heading := "Hello"
	if i%2 == 0 {
		heading = "Bonjour"
	}
	data := map[string]string{
		"title":   "Golang Web Server",
		"Heading": heading,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
