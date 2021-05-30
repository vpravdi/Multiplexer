package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//Read to the conn
	/*	bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bs)
	*/
	//Write to the conn
	fmt.Fprintf(conn, "hello can you hear me?")

}
