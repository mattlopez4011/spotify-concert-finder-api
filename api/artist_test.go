package main

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"golang.org/x/oauth2"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"testing"
//	"time"
//
//	"github.com/zmb3/spotify"
//	"go.uber.org/zap"
//	"golang.org/x/oauth2/clientcredentials"
//)
//
//const (
//	machineGunKellyID = "6TIYQ3jFPwQSRmorSezPxX"
//	accessToken       = ""
//	refreshToken      = ""
//)
//
//func Test_api_getArtist(t *testing.T) {
//
//	z, _ := zap.NewDevelopment()
//
//	// Spotify
//	config := &clientcredentials.Config{
//		ClientID:     os.Getenv("SPOTIFY_ID"),
//		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
//		TokenURL:     spotify.TokenURL,
//	}
//	//token, err := config.Token(context.Background())
//	//if err != nil {
//	//	log.Fatalf("couldn't get token: %v", err)
//	//}
//
//	token := &oauth2.Token{
//		AccessToken:  accessToken,
//
//	}
//
//	client := spotify.Authenticator{}.NewClient(token)
//
//	spotifyModel := SpotifyClient{
//		client:        client,
//		authenticator: spotify.Authenticator{},
//		clientID:      "",
//		clientSecret:  "",
//		state:         "",
//		codeVerifier:  "",
//		codeChallenge: "",
//		errors:        nil,
//		results:       nil,
//		currentUserID: "",
//	}
//
//	t.Run("Test GET Machine Gun Kelly ", func(t *testing.T) {
//
//		a := &api{
//			Log:     z.Sugar(),
//			Spotify: spotifyModel,
//		}
//
//		router := newRouter(a)
//
//		w := httptest.NewRecorder()
//		r := httptest.NewRequest(http.MethodGet, fmt.Sprintf(`/artist/%s`, machineGunKellyID), nil)
//
//		router.ServeHTTP(w, r)
//
//		assert.Equal(t, http.StatusOK, w.Result().StatusCode) // want 200
//
//	})
//
//}
