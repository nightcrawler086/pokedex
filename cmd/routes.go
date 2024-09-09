package main

import (
	"example/pokedex/internal/config"
	"html/template"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, config *config.Config, logger *slog.Logger, tmpl *template.Template) {
	mux.Handle("/", handleIndex(config, logger, tmpl))
}
