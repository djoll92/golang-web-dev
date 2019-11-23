package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(writer, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(writer, "in chrome go to: dev tools / application / cookies")
}

func read(writer http.ResponseWriter, request *http.Request) {
	c, err := request.Cookie("my-cookie")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(writer, "YOUR COOKIE:", c)
}
