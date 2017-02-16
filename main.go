package main

import (
	"github.com/pressly/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	http.ListenAndServe(":3000", r)
}
