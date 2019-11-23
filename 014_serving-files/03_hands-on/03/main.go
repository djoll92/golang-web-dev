package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", handleTemplates)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTemplates(writer http.ResponseWriter, request *http.Request) {
	err := tpl.Execute(writer, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
