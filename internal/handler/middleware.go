package handler

import (
	"context"
	"merch/pkg/jwt"
	"net/http"
	"strings"
)

type contextKey string

const userIdKey contextKey = "userId"

func (h *Handler) middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 && headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		userId, err := jwt.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
