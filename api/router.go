package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (api *api) newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
		api.Log.Info("Hey there")
	})
	r.Get("/callback", completeAuth)

	return r
}


