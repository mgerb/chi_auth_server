package middleware

import (
	"context"
	"net/http"
)

//
type Claims struct {
	Email string
}

// JWTMiddleware-
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "jwt", &Claims{
			Email: "test123",
		})))
	})
}
