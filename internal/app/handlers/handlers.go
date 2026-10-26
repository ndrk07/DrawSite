package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func SignInPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signin.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	fmt.Println("user in sign in page")

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	fmt.Println("user in sign up page")

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

// func handleUsersCreate() http.HandlerFunc {
// 	type request struct {
// 		Login    string
// 		Password string
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		req := &request{}
// 		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		u := &model.User{}
// 	}
// }
