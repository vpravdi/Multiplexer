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
		log.Println(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)
		}
		if ln == "" {
			break
		}
		i++
	}

}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en><head><metacharset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\r \n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\n")
	fmt.Fprintf(conn, "\n")
	fmt.Fprintf(conn, body)
}
