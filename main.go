package main

import (
	"drawserver/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// mux.HandleFunc("/postsignup", getSignUpData)

}
