package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

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

		// store on server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", s)

}
