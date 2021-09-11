package main

import (
	"net/http"
)

func (a *api) login(w http.ResponseWriter, r *http.Request) {
	a.zLogger.Info("Hello World!")
	url := createAuthUrl()
	// wait for auth to complete

	a.zLogger.Info(url)
}
