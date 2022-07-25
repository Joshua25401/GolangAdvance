package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Config) routes() http.Handler {
	// Create router
	mux := chi.NewRouter()

	// Setup middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// Define application routes
	mux.Get("/", app.Dashboard)
	mux.Get("/login", app.LoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Get("/activate-account", app.ActivateAccount)

	mux.Post("/register", app.PostRegisterPage)
	mux.Post("/login", app.PostLoginPage)

	return mux
}
