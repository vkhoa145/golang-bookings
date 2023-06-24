package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vkhoa145/golang-bookings/pkg/config"
	"github.com/vkhoa145/golang-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// pat libary
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/product", http.HandlerFunc(handlers.Repo.Product))
	// mux.Get("/order", http.HandlerFunc(handlers.Repo.Order))

	// chi libary
	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/product", handlers.Repo.Product)
	mux.Get("/order", handlers.Repo.Order)
	return mux
}
