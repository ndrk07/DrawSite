package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func signInPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user in sign in page")

	io.WriteString(w, "you in sign in page")
}

func signUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/start.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	fmt.Println("user in sign up page")

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}
