package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "https//*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Accept", "Content-Type", "X-CSRF_Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Get("/api/user/{id}", app.GetUserByID)

	mux.Get("/api/createUser", app.CreateUser)

	return mux
}
