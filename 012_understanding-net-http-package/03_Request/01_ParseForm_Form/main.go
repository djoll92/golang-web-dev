package main

import (
	"html/template"
	"log"
	"net/http"
)

var myTemplate *template.Template

type myHandler int

func (mh myHandler) ServeHTTP(myResponse http.ResponseWriter, myRequest *http.Request) {
	err := myRequest.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	myTemplate.ExecuteTemplate(myResponse, "index.gohtml", myRequest.Form)
}

func init() {
	myTemplate = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var myHandlerInstance myHandler
	http.ListenAndServe(":8080", myHandlerInstance)
}
