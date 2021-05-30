package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "\n TCP connection acknowledged")
		fmt.Fprintln(conn, "\nTCP count 1")
		fmt.Fprintf(conn, "\n%v", "TCP count 2")
		fmt.Fprintln(conn, "\nTerminating multiplexer")
		conn.Close()
	}
}
