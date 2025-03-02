package main

import (
	"github.com/bobomurod/go-auth-bolilerplate/internal/adapters/baseLogger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			return
		}
	})
	return r
}

func main() {
	baseLogger := baseLogger.NewSlogLogger()
	r := setupRouter()
	baseLogger.Info("Starting server")
	err := http.ListenAndServe(":3111", r)
	if err != nil {
		baseLogger.Error("Error starting server", "error", err)
		return
	}
}
