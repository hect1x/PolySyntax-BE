package middleware

import (
	"net/http"
	"strings"

	"polysyntax-be/pkg/response"
)

func APIKeyAuth(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token, found := strings.CutPrefix(authHeader, "Bearer ")
			if !found || token != apiKey {
				response.JSON(w, http.StatusUnauthorized, response.Response{Success: false, Error: "unauthorized"})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
