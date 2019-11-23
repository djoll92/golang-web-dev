package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is default route")
}

func dogFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Ruff ruff")
}

func meFunc(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Djordje")
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(defaultFunc))
	http.Handle("/dog/", http.HandlerFunc(dogFunc))
	http.Handle("/me/", http.HandlerFunc(meFunc))

	http.ListenAndServe(":8080", nil)

}
