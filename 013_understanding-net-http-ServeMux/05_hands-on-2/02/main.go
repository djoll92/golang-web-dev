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

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
		conn.Close()
	}

}
