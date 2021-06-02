package main

import (
	"fmt"
	"log"
	"net/http"
)

type trial int

func (m trial) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Request handled")
	fmt.Println(w, "Request")
}

func main() {
	var d trial
	http.ListenAndServe(":8080", d)
}
