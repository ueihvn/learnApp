package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/form.html"))
	t.Execute(w, "Hello World!	")
}
