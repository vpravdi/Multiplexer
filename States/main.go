package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", a)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func a(res http.ResponseWriter, req *http.Request) {
	a, b := req.FormValue("a"), req.FormValue("b")
	io.WriteString(res, "dog:"+a)
	io.WriteString(res, "cat:"+b)
}
