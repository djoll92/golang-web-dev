package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myFunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func myFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintln(writer, "go look at your terminal")
}
