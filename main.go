package main

import (
	"drawserver/internal/app/apiserver"
	"log"
)

var config = &apiserver.Config{
	BindAddr:    ":8080",
	DatabaseURL: "user=postgres password=an101978 host=localhost dbname=drawsite sslmode=disable",
}

func main() {
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
