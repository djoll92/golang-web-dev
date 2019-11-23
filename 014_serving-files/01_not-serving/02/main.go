package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}

func myFunc(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(writer, `
	<!--image doesn't serve'--!>
	<img src="/toby.jpg">
	`)

}
