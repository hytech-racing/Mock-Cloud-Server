package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	port := ":3000"
	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get route called")
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("upload route called")
	})

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}
