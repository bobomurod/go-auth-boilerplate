package http

import (
	"context"
	"fmt"
	"github.com/bobomurod/go-auth-bolilerplate/internal/config"
	"github.com/bobomurod/go-auth-bolilerplate/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Server struct {
	server *http.Server
	router *chi.Mux
	logger ports.Logger
}

func NewServer(cfg config.HTTPConfig, logger ports.Logger) *Server {
	router := chi.NewRouter()

	//be careful with middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(cfg.ReadTimeout))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &Server{
		server: server,
		router: router,
		logger: logger,
	}
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
