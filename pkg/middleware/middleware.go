package middleware

import (
	"cardValidator/pkg/responder"
	"log"
	"net/http"
	"runtime/debug"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				responder.RespondWithError(w, http.StatusInternalServerError, "occurred error")
				log.Println("Recovered from panic: %s", string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
