package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func abundance(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:  "general",
		Value: "some general cookie value",
	})
	http.SetCookie(writer, &http.Cookie{
		Name:  "specific",
		Value: "some specific cookie value",
	})
	fmt.Fprintln(writer, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(writer, "in chrome go to: dev tools / application / cookies")
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
	c1, err := request.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(writer, "YOUR COOKIE #1:", c1)
	}

	c2, err := request.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(writer, "YOUR COOKIE #2:", c2)
	}

	c3, err := request.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(writer, "YOUR COOKIE #3:", c3)
	}
}
