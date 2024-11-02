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

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func loadRoutes(s3Repo *S3Repository) *chi.Mux {
	router := chi.NewRouter()
	
	router.Use(middleware.Logger)

	router.Get("/api/v2/mcap/getNew", func(w http.ResponseWriter, r *http.Request) {
		var matchingEntries []DataEntryNew
		
		enableCORS(&w)

		w.WriteHeader(http.StatusOK)

		queryParams := r.URL.Query()

		file, err := os.Open("data/data-new.json")
		if err != nil {
			http.Error(w, "Failed to open data file", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer file.Close()

		matchingEntries = ParseJSONNew(file, queryParams)

		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		w.Header().Set("Content-Type", "application/json")
		res := make(map[string]interface{})
		res["data"] = matchingEntries
		err = encoder.Encode(res)
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
