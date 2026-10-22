package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", startHttp)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func startHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user in start page")

	io.WriteString(w, "Hello this is start page")
}
