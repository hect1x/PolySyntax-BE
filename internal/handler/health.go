package handler

import (
	"net/http"

	"polysyntax-be/pkg/response"

	"github.com/go-chi/chi/v5"
)

func RegisterHealthRoutes(r chi.Router) {
	r.Get("/health", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response.OK(w, map[string]string{"status": "ok"})
}
