package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *api) newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Get("/login", a.login)
	r.Get("/callback", a.completeAuth)

	return r
}

func (a *api) startServer() {
	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", a.config.port),
		Handler: a.newRouter(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	a.zLogger.Info("Starting on server port")

	//Need to return error
	
	err := srv.ListenAndServe()
	if err != nil {
		a.zLogger.Error(err)
	}
}


