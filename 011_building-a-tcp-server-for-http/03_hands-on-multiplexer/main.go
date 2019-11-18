package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	rl := strings.Fields(ln)
	m := rl[0]   // method
	uri := rl[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", uri)

	// multiplexer
	if m == "GET" && uri == "/" {
		index(conn)
	}
	if m == "GET" && uri == "/about" {
		about(conn)
	}
	if m == "GET" && uri == "/contact" {
		contact(conn)
	}
	if m == "GET" && uri == "/apply" {
		apply(conn)
	}
	if m == "POST" && uri == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Index</title>
</head>
<body>
<strong>INDEX</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>About</title>
</head>
<body>
<strong>ABOUT</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Contact</title>
</head>
<body>
<strong>CONTACT</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Apply</title>
</head>
<body>
<strong>APPLY</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
<form method="post" action="/apply">
<input type="submit" value="apply">
</form>
</body>
</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Apply Process</title>
</head>
<body>
<strong>APPLY PROCESS</strong><br>
<a href="/">index</a><br>
<a href="/about">about</a><br>
<a href="/contact">contact</a><br>
<a href="/apply">apply</a><br>
</body>
</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
