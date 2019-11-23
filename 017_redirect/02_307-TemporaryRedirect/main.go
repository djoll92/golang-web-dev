package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Your request method at foo: ", request.Method, "\n\n")
}

func bar(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Your request method at bar:", request.Method)
	// we could process form submission here
	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func barred(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Your request method at barred:", request.Method)
	tpl.ExecuteTemplate(writer, "index.gohtml", nil)
}
