package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {

	f := request.FormValue("first")
	l := request.FormValue("last")
	s := request.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(writer, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(writer, err.Error(), 500)
		log.Fatalln(err)
	}
}
