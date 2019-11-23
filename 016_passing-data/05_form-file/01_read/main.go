package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(writer http.ResponseWriter, request *http.Request) {

	var s string
	fmt.Println(request.Method)
	if request.Method == http.MethodPost {
		//open
		f, h, err := request.FormFile("q")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// FYI
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, `
	<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	<br>`+s)

}
