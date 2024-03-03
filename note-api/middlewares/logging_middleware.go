package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogginHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v\n", r.URL.Path, time.Since(start))
	})
}
