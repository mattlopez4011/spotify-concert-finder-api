package main

import (
	"fmt"
	"net/http"
	"spotify-concert-finder-api/models"

	"github.com/go-chi/chi"
	"github.com/zmb3/spotify"
)

func (a *api) getArtist(w http.ResponseWriter, r *http.Request) {
	a.Log.Infof("🎸 Starting getArtist...")
	a.Log.Infof("GetArtist request: %+v", r)

	artistID := chi.URLParam(r, "artist_id")

	a.Log.Infof("Artist ID from params: %s", artistID)

	artist := models.Artist{}

	err := artist.GetArtistByID(a.Log, artistID)
	if err != nil {
		a.Log.Errorw("Error getting artist by ID", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	id := spotify.ID("").String()

	a.Log.Infof("🎵 a.Spotify: %+v", a.Spotify)
	a.Log.Infof("a.Spotify.authenticator: %+v", a.Spotify.authenticator)

	artist, err := a.Spotify.client.GetArtist(spotify.ID(artistID))
	if err != nil {
		a.Log.Errorw("Error getting artist by ID", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	a.Log.Infof("Artist: %+v", artist)

	_, _ = w.Write(jsonBody)
	w.WriteHeader(http.StatusOK)

	fmt.Println()
	fmt.Println()
}
