package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New(s3Repo *S3Repository) *App { // Accept *S3Repository as a parameter
	app := &App{
		router: loadRoutes(s3Repo), // Pass s3Repo to loadRoutes
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("failed to start server %s", err)
	}

	return nil
}
