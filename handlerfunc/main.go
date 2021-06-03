package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/b/", b)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)

}

func a(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "a ran")
}

func b(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "b ran")
}

func myName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello Praveen")
}
