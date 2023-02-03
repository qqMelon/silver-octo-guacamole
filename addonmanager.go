package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("template/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Valeur: "Hello World"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
