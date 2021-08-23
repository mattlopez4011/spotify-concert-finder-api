package main

import (
	"go.uber.org/zap"
	"testing"
)

func Test_newRouter(t *testing.T) {

	t.Run("Test router", func(t *testing.T) {
		z, _ := zap.NewDevelopment()

		a := &api{
			Log:     z.Sugar(),
			Spotify: SpotifyClient{},
		}

		newRouter(a)

	})

}
