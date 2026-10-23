package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/signup", signUpPage)
	mux.HandleFunc("/postsignup", getSignUpData)
	mux.HandleFunc("/signin", signInPage)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
