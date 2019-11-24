package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(writer, cookie)
	}
	fmt.Println(cookie)
}
