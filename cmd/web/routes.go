package main

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(a.LoadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		a.session.Put(r.Context(), "test", "Tobi Atilola")
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		vars := make(jet.VarMap)
		vars.Set("test", a.session.GetString(r.Context(), "test"))
		err := a.render(w, r, "index", vars)
		if err != nil {
			log.Fatal(err)
		}
	})

	return mux
}
