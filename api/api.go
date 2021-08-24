package main

import (

	_ "github.com/lib/pq"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
)

type api struct {
	Log     *zap.SugaredLogger
// 	Spotify SpotifyClient
}

type SpotifyClient struct {
	client        spotify.Client
	authenticator spotify.Authenticator
	clientID      string
	clientSecret  string
	state         string
	codeVerifier  string
	codeChallenge string
	errors        chan error
	results       chan string
	currentUserID string
}

func newProductionApi() (*api, error) {
	// zap for logging
	z, _ := zap.NewProduction()

	// Spotify
	//config := &clientcredentials.Config{
	//	ClientID:     os.Getenv("SPOTIFY_ID"),
	//	ClientSecret: os.Getenv("SPOTIFY_SECRET"),
	//	TokenURL:     spotify.TokenURL,
	//}
	//token, err := config.Token(context.Background())
	//if err != nil {
	//	log.Fatalf("couldn't get token: %v", err)
	//}
	//
	//token := &oauth2.Token{
	//			AccessToken:  "BQDaSac-zFzd_uH75D8DiiXUwSrTnvGqCfrZeeIKu09y-9M7vNzIiEKahkM3PBhtwIbyOkwor8sSlhyXAhcKdulrIA9Hv6ndujtmCt2dZY8miiORW1yhAfKajf3TT00pOzup99pzwTqlZ7D-BQNMcygERmR_mbvbjfAeJiUopbsinPlvzqpn6CbNmNQ07XGP_vOeQEHHssf-8Y0R10LAOEw6DiZuFCxt1wcol5S-3YLKcJ-2KZuoyA9HJFfn6w",
	//}
// 	tokenFromDB := "BQDaSac-zFzd_uH75D8DiiXUwSrTnvGqCfrZeeIKu09y-9M7vNzIiEKahkM3PBhtwIbyOkwor8sSlhyXAhcKdulrIA9Hv6ndujtmCt2dZY8miiORW1yhAfKajf3TT00pOzup99pzwTqlZ7D-BQNMcygERmR_mbvbjfAeJiUopbsinPlvzqpn6CbNmNQ07XGP_vOeQEHHssf-8Y0R10LAOEw6DiZuFCxt1wcol5S-3YLKcJ-2KZuoyA9HJFfn6w"
//
// 	token := new(oauth2.Token)
//
// 	err := json.Unmarshal([]byte(tokenFromDB), token)
// 	if err != nil {
// 		fmt.Println("error unmarshalling token from DB ", err.Error())
// 	}

// 	client := spotify.NewAuthenticator("").NewClient(token)

	//client := spotify.Authenticator{}.NewClient(token)

// 	spotifyModel := SpotifyClient{
// 		client: client,
// 		//client:        spotify.Client{},
// 		authenticator: spotify.Authenticator{},
// 		clientID:      os.Getenv("SPOTIFY_ID"),
// 		clientSecret:  os.Getenv("SPOTIFY_SECRET"),
// 		state:         "",
// 		codeVerifier:  "",
// 		codeChallenge: "",
// 		errors:        nil,
// 		results:       nil,
// 		currentUserID: "",
// 	}
//
// 	spotifyModel.authenticator.SetAuthInfo(os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))

	a := &api{
		Log:     z.Sugar(),
// 		Spotify: spotifyModel,
	}

	return a, nil
}
