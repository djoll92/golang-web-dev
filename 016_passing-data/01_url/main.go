package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {
	v := request.FormValue("q")
	io.WriteString(writer, "Do my search: "+v)
}

// visit this page:
// localhost:8080/?q=dog
