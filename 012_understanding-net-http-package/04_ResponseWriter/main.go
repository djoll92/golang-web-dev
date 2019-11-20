package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var handlerObject myHandler
	http.ListenAndServe(":8080", handlerObject)
}
