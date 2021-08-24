package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func newRouter(a *api) *chi.Mux {
	if a == nil {
		api, err := newProductionApi()
		if err != nil {
			panic(err.Error())
		}

		a = api
	}

	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) })
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusMethodNotAllowed) })

// 	r.Get("/artist/{artist_id}", a.getArtist)
// 	r.Get("/login-spotify", a.startSpotifyAuthentication)
// 	r.Get("/login-spotify/successful", a.spotifyAuthenticationSuccessful)

	//Test Endpoint
	r.Get("/hello-world", a.helloWorld)

	//err :=  http.ListenAndServe(":3000", r)
	//if err != nil {
	//	a.Log.Errorw("Error serving request", "error", err.Error())
	//}

	return r
}
