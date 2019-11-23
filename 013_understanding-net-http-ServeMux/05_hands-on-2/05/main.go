package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	io.WriteString(conn, "I see you've connected.")
}
