package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserFollowRead), spotifyauth.WithClientID("1ea7b85e0d5644e1bf288a172bb5445f"), spotifyauth.WithClientSecret("256e00f442cd4f13a22b208a12d30506"))
	ch    = make(chan *spotify.Client)
	state = "abc123"
	// These should be randomly generated for each request
	//  More information on generating these can be found here,
	// https://www.oauth.com/playground/authorization-code-with-pkce.html
	codeVerifier  = "w0HfYrKnG8AihqYHA9_XUPTIcqEXQvCQfOF2IitRgmlF43YWJ8dy2b49ZUwVUOR.YnvzVoTBL57BwIhM4ouSa~tdf0eE_OmiMC_ESCcVOe7maSLIk9IOdBhRstAxjCl7"
	codeChallenge = "ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig"
)

func main() {
	z, _ := zap.NewProduction()
	zLog := z.Sugar()
	r := chi.NewRouter()
	r.Use(middleware.Logger) 
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
		zLog.Info("Hello from z logger")
	})
	err := http.ListenAndServe(":8080", r)


	if err != nil {
		// a.Log.Errorw("Error serving request", "error", err.Error())
		log.Println(err.Error())
	}

	url := auth.AuthURL(state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)

	log.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	
	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("You are logged in as:", user.ID)

}
