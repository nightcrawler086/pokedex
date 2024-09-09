package main

import (
	"database/sql"
	"example/pokedex/internal/config"
	"html/template"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, config *config.Config, logger *slog.Logger, tmpl *template.Template, db *sql.DB) {
	mux.Handle("/", handleIndex(config, logger, tmpl, db))
}
