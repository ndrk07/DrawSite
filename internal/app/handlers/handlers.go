package handlers

import (
	"html/template"
	"io"
	"net/http"
)

func SignInPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signin.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func HandleMainMenu(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "succesfully registration")
}

func ErrorSignUp(w http.ResponseWriter, r *http.Request, err error) {
	tmpl, _ := template.ParseFiles("templates/signup.html")
	tmpl.Execute(w, err)
}
