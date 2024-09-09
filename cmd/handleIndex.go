package main

import (
	"database/sql"
	"encoding/json"
	"example/pokedex/internal/config"
	"html/template"
	"io"
	"log/slog"
	"net/http"
)

func handleIndex(config *config.Config, logger *slog.Logger, tmpl *template.Template, db *sql.DB) http.Handler {

	const pokemonUri = "http://pokeapi.co/api/v2/pokemon"
	type Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type Response struct {
		Name    string    `json:"name"`
		Results []Pokemon `json:"results"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Made it to Index Handler")
			logger.Info(config.Host)
			if err := db.Ping(); err != nil {
				logger.Error(err.Error())
			}
			response, err := http.Get(pokemonUri + "?limit=50")
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
		},
	)
}
