package main

import (
	"context"
	"database/sql"
	"example/pokedex/internal/config"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func NewServer(config *config.Config, logger *slog.Logger, tmpl *template.Template, db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, config, logger, tmpl, db)
	var handler http.Handler = mux
	return handler
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// Config
	config := config.NewConfig()
	// Templater
	tmpl := template.Must(template.ParseGlob("./views/*.html"))
	// DB
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "postgres", "5432", "postgres", "postgres", "pokedex")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error(err.Error())
	}
	srv := NewServer(config, logger, tmpl, db)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup // need to understand this better
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error running server: %s\n", err)
		os.Exit(1)
	}
}
