package handlers

import (
	"html/template"
	"io"
	"net/http"
)

func HandleMainMenu(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "succesfully registration")
}

func ErrorOutputTmpl(w http.ResponseWriter, r *http.Request, tmpl *template.Template, err error) {
	tmpl.Execute(w, err)
}
