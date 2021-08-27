package main

// import (
// 	"testing"
// )

// func Test_api_startSpotifyAuthentication(t *testing.T) {
// 	z, _ := zap.NewDevelopment()

// 	t.Run("Test Send user to Spotify client for user token", func(t *testing.T) {
// 		a := &api{
// 			Log:     z.Sugar(),
// 			Spotify: SpotifyClient{},
// 		}
	
// 		router := newRouter(a)
	
// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest(http.MethodGet, `/login-spotify`, nil)
	
// 		//router.ServeHTTP(w, r)
	
// 		assert.Equal(t, http.StatusOK, w.Result().StatusCode) // want 200
// 	})

// }
