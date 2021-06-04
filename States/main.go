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
	a := req.FormValue("a")
	//io.WriteString(res, "dog:\t"+a+"\n")
	//io.WriteString(res, "cat:\t"+b+"\n")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="post">
		<input type="text" name="a">
		<input type="submit">
	</form>
	<br>`+a)
}
