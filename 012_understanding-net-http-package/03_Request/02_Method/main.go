package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}

	myTemplate.ExecuteTemplate(w, "index.gohtml", data)
}

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var handlerInstance myHandler
	http.ListenAndServe(":8080", handlerInstance)
}
