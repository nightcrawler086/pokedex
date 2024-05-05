package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

const pokemonUri = "http://pokeapi.co/api/v2/pokemon"

var tmpl *template.Template

type Response struct {
	Name    string    `json:"name"`
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func init() {
	if tmpl == nil {
		tmpl = template.Must(template.ParseGlob("views/*.html"))
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(pokemonUri + "?limit=20")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseObj Response
	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		panic(err)
	}
	data := map[string]any{
		"title":         "Pokedex",
		"Heading":       "Pokedex",
		"queryResponse": responseObj,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
