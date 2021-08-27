package main

import (
	"testing"

	"go.uber.org/zap"
)

func Test_newRouter(t *testing.T) {

	t.Run("Test router", func(t *testing.T) {
		z, _ := zap.NewDevelopment()

		a := &api{
			Log:     z.Sugar(),
			//Spotify: SpotifyClient{},
		}

		newRouter(a)

	})

}
