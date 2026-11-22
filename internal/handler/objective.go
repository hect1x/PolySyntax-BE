package handler

import (
	"net/http"

	"polysyntax-be/internal/domain"
	"polysyntax-be/pkg/response"

	"github.com/go-chi/chi/v5"
)

func RegisterObjectiveRoutes(r chi.Router) {
	r.Get("/objective/example", getObjectiveExample)
}

func getObjectiveExample(w http.ResponseWriter, r *http.Request) {
	example := domain.Objective{
		ID:    1,
		Title: "Arrow Functions",
		Snippets: []domain.Snippet{
			{Language: "JavaScript", Code: "const add = (a, b) => a + b"},
			{Language: "TypeScript", Code: "const add = (a: number, b: number): number => a + b"},
		},
	}

	response.OK(w, example)
}
