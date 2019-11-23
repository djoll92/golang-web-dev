package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(writer, "file not found", 404)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(writer, "file not found", 404)
		return
	}

	http.ServeContent(writer, request, file.Name(), fileInfo.ModTime(), file)
}

func dog(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, `<img src="toby.jpg">`)
}
