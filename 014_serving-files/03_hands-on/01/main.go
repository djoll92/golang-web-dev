package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", servePic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func servePic(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "toby.jpg")
}

func dog(writer http.ResponseWriter, request *http.Request) {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
	err := tpl.ExecuteTemplate(writer, "dog.gohtml", nil)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
}

func foo(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "foo ran")
}
