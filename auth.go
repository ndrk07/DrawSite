package main

import (
	"fmt"
	"net/http"
)

func getSignUpData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("pw")
	repeatPassword := r.FormValue("pw_ag")

	fmt.Println(name, password, repeatPassword)
}
