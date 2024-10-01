package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	var matchingEntries []DataEntry

	router.Use(middleware.Logger)
	router.Get("/api/v2/mcap", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		queryParams := r.URL.Query()

		// date := queryParams.Get("date")
		// location := queryParams.Get("location")
		// notes := queryParams.Get("notes")
		// eventType := queryParams.Get("eventType")

		// response := fmt.Sprintf("date: %s, location: %s, notes: %s, eventType: %s", date, location, notes, eventType)
		// fmt.Println(response)

		file, err := os.Open("data/data.json")

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		matchingEntries = ParseJSON(file, queryParams)
		fmt.Println(matchingEntries)

	})

	return router
}
