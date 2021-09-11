package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *api) newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Get("/", a.login)
	r.Get("/callback", completeAuth)

	return r
}


