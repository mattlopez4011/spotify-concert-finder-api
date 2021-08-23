package main

import (
	"encoding/json"
	"net/http"
)

type RequestBody struct {
	User  User   `json:"user"`
	Notes string `json:"notes"`
}

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *api) helloWorld(w http.ResponseWriter, r *http.Request) {
	a.Log.Info("Starting HelloWorld...")
	a.Log.Infof("HelloWorld request: %+v", r)

	//Init request body
	requestBody := RequestBody{User: User{
		UserID:    "1234asdf",
		FirstName: "John",
		LastName:  "Smith",
	}}

	requestBody.Notes = "Testing Notes from backend"

	//Send json via request body
	encoder := json.NewEncoder(w)
	err := encoder.Encode(requestBody)
	if err != nil {
		a.Log.Errorw("Json encode error", "error", err.Error(), "userID", requestBody.User.UserID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Set header status code to 200
}
