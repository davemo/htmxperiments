package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.Use(loggingMiddleware)

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}