package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "toby.jpg")
}

func dog(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, `<img src="toby.jpg">`)
}
