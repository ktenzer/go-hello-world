package main

import (
	"log"
	"net/http"
	"time"
)

func LogApi(handler http.Handler, apiRoute string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		handler.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			apiRoute,
			time.Since(startTime),
		)
	})
}
