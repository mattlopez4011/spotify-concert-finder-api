package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
//
// 	"github.com/zmb3/spotify"
// 	"spotify-concert-finder-api/models"
// )
//
// const redirectURI = "http://192.168.2.4:3000/login-spotify/successful"
//
// //const redirectURI = "http://localhost:3000/login-spotify/successful"
// // http://localhost:3000/login-spotify
// // redirectURI = "http://127.0.0.1:8080/callback"
//
// func (a *api) startSpotifyAuthentication(w http.ResponseWriter, r *http.Request) {
// 	a.Log.Infof("1️⃣ Starting startSpotifyAuthentication...")
// 	a.Log.Infof("startAuth request:%+v", r)
//
// 	a.Spotify.authenticator = spotify.NewAuthenticator(
// 		redirectURI,
// 		spotify.ScopeUserFollowRead,
// 		spotify.ScopeUserReadPrivate,
// 		spotify.ScopePlaylistReadPrivate,
// 		spotify.ScopePlaylistModifyPublic,
// 		spotify.ScopePlaylistModifyPrivate,
// 	)
//
// 	a.Spotify.authenticator.SetAuthInfo(os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))
//
// 	state, err := models.GenerateRandomString()
// 	if err != nil {
// 		a.Log.Errorw("Error generating state", "error", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, _ = w.Write([]byte(err.Error()))
// 	}
//
// 	a.Spotify.state = state
// 	a.Log.Infof("state: %s", state)
//
// 	url := a.Spotify.authenticator.AuthURL(state)
// 	a.Log.Infof("auth URL: %+v", url)
//
// 	//Need to send this url to front end
// 	//Front end will need to redirect user to provided url
// 	//err := browser.OpenURL(url)
// 	//if err != nil {
// 	//	a.Log.Errorw("Error opening browser")
// 	//}
//
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Println()
// 	fmt.Println()
// }
//
// func (a *api) spotifyAuthenticationSuccessful(w http.ResponseWriter, r *http.Request) {
// 	a.Log.Infof("2️⃣ Starting spotifyAuthenticationSuccessful...")
// 	a.Log.Infof("spotifyAuthenticationSuccessful request:%+v", r)
//
// 	errors, ok := r.URL.Query()["error"]
// 	if ok {
// 		fmt.Fprintf(w, "<html><body><p><strong>Error</strong>: %s</p></body></html>", errors[0])
// 		a.Spotify.errors <- fmt.Errorf("error: %s", errors[0])
//
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, _ = w.Write([]byte(errors[0]))
// 		return
// 	}
//
// 	a.Log.Infof("State before token is requested: %+v", a.Spotify.state)
//
// 	token, err := a.Spotify.authenticator.Token(a.Spotify.state, r)
// 	if err != nil {
// 		a.Log.Errorw("Error requesting access token", "error", err.Error())
//
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, _ = w.Write([]byte(err.Error()))
// 		return
// 	}
// 	a.Log.Infof("Token: %+v", token)
// 	a.Log.Infof("token.AccessToken: %+v", token.AccessToken)
// 	a.Log.Infof("token.RefreshToken: %+v", token.RefreshToken)
//
// 	client := a.Spotify.authenticator.NewClient(token)
//
// 	user, err := client.CurrentUser()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, _ = w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	a.Spotify.client = client
// 	a.Spotify.currentUserID = user.ID
//
// 	a.Log.Infof("a.Spotify: %+v", a.Spotify)
// 	a.Log.Info("========================================================")
// 	a.Log.Infof("Current user: %+v", user)
// 	a.Log.Info("========================================================>>>>>>>>>>>>>>>>")
//
// 	followedArtist, err := client.CurrentUsersFollowedArtistsOpt(50, "")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		_, _ = w.Write([]byte(err.Error()))
// 		return
// 	}
//
// 	a.Log.Infof("followed artist: %+v", followedArtist)
// 	a.Log.Infof("Artist array: %+v", followedArtist.Artists)
//
// 	// Respond
// 	encoder := json.NewEncoder(w)
// 	err = encoder.Encode(followedArtist)
// 	if err != nil {
// 		a.Log.Errorw("Json encode error", "error", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.WriteHeader(http.StatusOK)
//
// 	fmt.Println()
// 	fmt.Println()
// }


