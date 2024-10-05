package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
			fmt.Println(err)
		}

		defer file.Close()

		matchingEntries = ParseJSON(file, queryParams)
		// fmt.Println(matchingEntries)

		for _, entry := range matchingEntries {
			fmt.Println("BUCKET: " + entry.Bucket)
			fmt.Println("PATH: " + entry.Path)

			signedUrl := s3Repo.GetSignedUrl(r.Context(), entry.Bucket, entry.Path)
			fmt.Println(signedUrl)
		}

	})

	return router
}
