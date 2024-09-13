package main

import (
	"github.com/chronalx/url-shortener/internal/app/handlers"
	_ "github.com/chronalx/url-shortener/internal/app/storage"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRequests)
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
