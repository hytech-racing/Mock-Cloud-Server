package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hytech-racing/Mock-Cloud-Server/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter();

	router.Use(middleware.Logger);
	router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/api/v2/mcap/", loadFileRoutes)

	return router
}

func loadFileRoutes(router chi.Router) {
	fileHandler := &handler.File{}

	router.Post("/upload", fileHandler.UploadFile)
}