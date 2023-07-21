package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) Writer(w http.ResponseWriter, r *http.Request, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return w

}

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Get("/", app.Home)

	return mux
}
