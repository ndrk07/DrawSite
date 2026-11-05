package main

import (
	"drawserver/internal/app/apiserver"
	"log"
)

func main() {
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
