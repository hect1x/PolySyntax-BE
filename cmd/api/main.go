package main

import (
	"log"
	"net/http"

	"polysyntax-be/config"
	"polysyntax-be/internal/handler"
	appmiddleware "polysyntax-be/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		handler.RegisterHealthRoutes(r)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(appmiddleware.APIKeyAuth(cfg.APIKey))
			// handler.RegisterXxxRoutes(r)
		})
	})

	log.Printf("Server running on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
