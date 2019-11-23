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
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", templateHandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func templateHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	err := tpl.Execute(writer, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
