package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

func LogginHandler(next http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s\n", r.Method, r.URL.Path)
		// log to file
		handlers.LoggingHandler(logFile, next)
		log.Printf("Completed %s in %v\n", r.URL.Path, time.Since(start))
	})
}
