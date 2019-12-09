package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "admin:djole123@tcp(mydbinstance.c2lbv4s7p3za.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8")
	check(err)
	defer db.Close()

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(writer http.ResponseWriter, request *http.Request) {
	_, err = io.WriteString(writer, "Successfully connected.")
	check(err)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
