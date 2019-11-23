package main

import (
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
		io.WriteString(conn, "I see you connected\n")
		conn.Close()
	}

}
