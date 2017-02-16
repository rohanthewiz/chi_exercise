package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home page"))
	})
	r.Get("/ping", func(w http.ResponseWriter, r * http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r * http.Request) {
		panic("Testing panic")
	})


	http.ListenAndServe(":3000", r)
}
