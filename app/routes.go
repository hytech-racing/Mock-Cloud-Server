package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hytech-racing/Mock-Cloud-Server/handler"
)

func loadRoutes(s3Repo *S3Repository) *chi.Mux {
	router := chi.NewRouter()
	var matchingEntries []DataEntry

	router.Use(middleware.Logger)
	router.Get("/api/v2/mcap", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		queryParams := r.URL.Query()

		file, err := os.Open("data/data.json")
		if err != nil {
			http.Error(w, "Failed to open data file", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer file.Close()

		matchingEntries = ParseJSON(file, queryParams)
		fmt.Println(matchingEntries)

		var entries []DataEntry
		for _, entry := range matchingEntries {
			signedUrl := s3Repo.GetSignedUrl(r.Context(), entry.Bucket, entry.Path)
			entry.SignedURL = signedUrl
			entries = append(entries, entry)
		}

		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		w.Header().Set("Content-Type", "application/json")
		err = encoder.Encode(entries)

		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	})

	router.Route("/api/v2/mcap/", loadFileRoutes)

	return router
}

func loadFileRoutes(router chi.Router) {
	fileHandler := &handler.File{}
	router.Post("/upload", fileHandler.UploadFile)
}
