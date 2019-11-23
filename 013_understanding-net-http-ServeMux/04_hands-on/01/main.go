package main

import (
	"io"
	"net/http"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is default route")
}

func dogFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Ruff ruff")
}

func meFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Djordje")
}

func main() {

	http.HandleFunc("/", defaultFunc)
	http.HandleFunc("/dog/", dogFunc)
	http.HandleFunc("/me/", meFunc)

	http.ListenAndServe(":8080", nil)

}
