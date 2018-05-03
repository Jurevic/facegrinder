package middleware

import (
	"net/http"
	"log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Access from remote address: %s, method: %s, url: %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}