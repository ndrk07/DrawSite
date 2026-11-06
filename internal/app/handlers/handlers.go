package handlers

import (
	"html/template"
	"net/http"
)

func ErrorOutputTmpl(w http.ResponseWriter, r *http.Request, tmpl *template.Template, err error) {
	tmpl.Execute(w, err)
}
