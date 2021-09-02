package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func newDevRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	r.Get("/callback", completeAuth)

	return r
}


