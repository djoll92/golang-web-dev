package main

import (
	"io"
	"net/http"
)

type router int

func (routerObject router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var routerInstance router
	http.ListenAndServe(":8080", routerInstance)
}
