package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	// "github.com/go-chi/chi"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// 	"github.com/zmb3/spotify/v2"
// 	spotifyauth "github.com/zmb3/spotify/v2/auth"
// 	"golang.org/x/oauth2"
// )

// const redirectURI = "http://localhost:8080/callback"

// var (
// 	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
// 	ch    = make(chan *spotify.Client)
// 	state = "abc123"
// 	// These should be randomly generated for each request
// 	//  More information on generating these can be found here,
// 	// https://www.oauth.com/playground/authorization-code-with-pkce.html
// 	codeVerifier  = "w0HfYrKnG8AihqYHA9_XUPTIcqEXQvCQfOF2IitRgmlF43YWJ8dy2b49ZUwVUOR.YnvzVoTBL57BwIhM4ouSa~tdf0eE_OmiMC_ESCcVOe7maSLIk9IOdBhRstAxjCl7"
// 	codeChallenge = "ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig"
// )

// type config struct {
// 	port int
// 	env  string
// }

// type AppStatus struct {
// 	Status      string `json:"status"`
// 	Environment string `json:"environment"`
// 	Version     string `json:"version"`
// }

// func main() {
// 	var cfg config
// 	os.Setenv("SPOTIFY_ID", "1ea7b85e0d5644e1bf288a172bb5445f")
// 	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
// 	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production ")
// 	flag.Parse()

// 	if cfg.env == "prod" {
// 		// lambda.Start(handler)
// 	} else {
// 		newDevRouter()

// 		url := auth.AuthURL(state,
// 			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
// 			oauth2.SetAuthURLParam("code_challenge", codeChallenge),
// 		)
// 		fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

// 		// wait for auth to complete
// 		client := <-ch

// 		// use the client to make calls that require authorization
// 		user, err := client.CurrentUser(context.Background())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("You are logged in as:", user.ID)
// 	}

// }

// // func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// // 	chiLambda := chiAdapter.New(newRouter(nil))

// // 	response, err := chiLambda.ProxyWithContext(ctx, req)
// // 	if err != nil {
// // 		return response, err
// // 	}

// // 	response.Headers = map[string]string{
// // 		"x-custom-header":              "*",
// // 		"Access-Control-Allow-Origin":  "*",
// // 		"Access-Control-Allow-Methods": "GET,HEAD,POST,OPTIONS",
// // 		"Access-Control-Max-Age":       "86400",
// // 	}

// // 	return response, nil
// // }

// func completeAuth(w http.ResponseWriter, r *http.Request) {
// 	tok, err := auth.Token(r.Context(), state, r,
// 		oauth2.SetAuthURLParam("code_verifier", codeVerifier))
// 	if err != nil {
// 		http.Error(w, "Couldn't get token", http.StatusForbidden)
// 		log.Fatal(err)
// 	}
// 	if st := r.FormValue("state"); st != state {
// 		http.NotFound(w, r)
// 		log.Fatalf("State mismatch: %s != %s\n", st, state)
// 	}
// 	// use the token to get an authenticated client
// 	client := spotify.New(auth.Client(r.Context(), tok))
// 	fmt.Fprintf(w, "Login Completed!")
// 	ch <- client
// }

// func newDevRouter() *chi.Mux {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Got request for:", r.URL.String())
// 	})
// 	r.Get("/callback", completeAuth)

// 	err := http.ListenAndServe(":8080", r)
// 	if err != nil {
// 		// a.Log.Errorw("Error serving request", "error", err.Error())
// 	}

// 	return r
// }


//IF YOU WANT THIS TO WORK
// 1) Register an app on spotify
// 2) Insert your client id and secret in the var auth
// 3) Run go run main.go
// 4) Paste url from terminal in your browser



import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2"

	"github.com/zmb3/spotify/v2"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate), spotifyauth.WithClientID("INSERT YOUR CLIENT ID HERE"), spotifyauth.WithClientSecret("INSERT YOUR SECRET HERE"))
	ch    = make(chan *spotify.Client)
	state = "abc123"
	// These should be randomly generated for each request
	//  More information on generating these can be found here,
	// https://www.oauth.com/playground/authorization-code-with-pkce.html
	codeVerifier  = "w0HfYrKnG8AihqYHA9_XUPTIcqEXQvCQfOF2IitRgmlF43YWJ8dy2b49ZUwVUOR.YnvzVoTBL57BwIhM4ouSa~tdf0eE_OmiMC_ESCcVOe7maSLIk9IOdBhRstAxjCl7"
	codeChallenge = "ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig"
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}
//https://accounts.spotify.com/authorize?client_id=1ea7b85e0d5644e1bf288a172bb5445f&code_challenge=ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig&code_challenge_method=S256&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&response_type=code&scope=user-read-private&state=abc123
// https://accounts.spotify.com/authorize?client_id=1ea7b85e0d5644e1bf288a172bb5445f&code_challenge=ZhZJzPQXYBMjH8FlGAdYK5AndohLzFfZT-8J7biT7ig&code_challenge_method=S256&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&response_type=code&scope=user-read-private&state=abc123
