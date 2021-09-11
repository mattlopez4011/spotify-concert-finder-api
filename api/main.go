package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	api, err := newApplication()
	if err != nil {
		log.Println(err.Error())
	}

	r := api.newRouter()

	http.ListenAndServe(":8080", r)
	

	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("You are logged in as:", user.ID)
}
