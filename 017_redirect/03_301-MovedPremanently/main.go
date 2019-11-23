package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Your request method at foo: ", request.Method, "\n\n")
}

func bar(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Your request method at bar:", request.Method)
	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}
