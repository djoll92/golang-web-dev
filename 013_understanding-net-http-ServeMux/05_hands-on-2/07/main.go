package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(nil)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(nil)
		}
		go serve(conn)
	}

}

func serve(conn net.Conn) {
	defer conn.Close()
	i := 0
	var method, uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		xs := strings.Fields(ln)
		if i == 0 {
			method = xs[0]
			uri = xs[1]
			fmt.Println("METHOD:", method)
			fmt.Println("URL:", uri)
		}
		if ln == "" {
			break
		}
		i++
	}
	body := "METHOD: " + method + "\nURL: " + uri
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
