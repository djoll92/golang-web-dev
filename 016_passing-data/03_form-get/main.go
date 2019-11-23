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
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}
