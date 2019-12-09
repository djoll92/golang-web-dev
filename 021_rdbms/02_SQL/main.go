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

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func del(writer http.ResponseWriter, request *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(writer, "DELETED RECORD", n)
}

func drop(writer http.ResponseWriter, request *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(writer, "DROPPED TABLE customer")

}

func update(writer http.ResponseWriter, request *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(writer, "UPDATED RECORD", n)
}

func read(writer http.ResponseWriter, request *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(writer, "RETRIEVED RECORD:", name)
	}
}

func insert(writer http.ResponseWriter, request *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(writer, "INSERTED RECORD", n)
}

func create(writer http.ResponseWriter, request *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(writer, "CREATED TABLE customer", n)

}

func amigos(writer http.ResponseWriter, request *http.Request) {
	rows, err := db.Query(`SELECT aName from amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(writer, s)
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
