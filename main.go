package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("./templates/css/"))))

	mux.HandleFunc("/", startHttp)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func startHttp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/start.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	fmt.Println("user in start page")

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}
