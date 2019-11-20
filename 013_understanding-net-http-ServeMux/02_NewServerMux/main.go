package main

import (
	"io"
	"net/http"
)

type dogHandler int

func (dh dogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "good boy!")
}

type catHandler int

func (ch catHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello kitty")
}

func main() {
	var d dogHandler
	var c catHandler

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
