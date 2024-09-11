package driving

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	user "github.com/kevin07696/ecommerce/domain/auth/services"
)

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for _, m := range middleware {
			next = m(next)
		}

		return next.ServeHTTP
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s. path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequireTokenMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func RequireSessionMiddleware(next http.Handler, userAPI user.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := userAPI.UpdateSession(r.Context(), w, r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func RequestIDMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request_id", uuid.New().String())

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
