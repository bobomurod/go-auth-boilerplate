package main

import (
	"context"
	"github.com/bobomurod/go-auth-bolilerplate/internal/adapters/baseLogger"
	repository "github.com/bobomurod/go-auth-bolilerplate/internal/adapters/repository/mongodb"
	"github.com/bobomurod/go-auth-bolilerplate/internal/config"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/services"
	"github.com/bobomurod/go-auth-bolilerplate/internal/framework/mongodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
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
	log := baseLogger.SlogLogger{}
	cfg := config.NewConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mongoClient, err := mongodb.NewMongoClient(ctx, mongodb.MongoConfig{
		URI:            cfg.MongoDB.URI,
		Database:       cfg.MongoDB.Database,
		ConnectTimeout: cfg.MongoDB.ConnectTimeout,
	})
	if err != nil {
		log.Error("Error connecting to mongodb", "error", err)
		os.Exit(1)
	}
	r := setupRouter()
	err = http.ListenAndServe(":3111", r)
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
	defer mongoClient.Close(ctx)
	userRepo := repository.NewUserRepository(mongoClient.Client())
	userService := services.NewUserService(userRepo, &log)
	log.Info("Starting server successfully")
}
