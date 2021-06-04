package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", a)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func a(res http.ResponseWriter, req *http.Request) {
	var a string
	fmt.Println(req.Method)

	//io.WriteString(res, "dog:\t"+a+"\n")
	//io.WriteString(res, "cat:\t"+b+"\n")
	if req.Method == http.MethodPost {
		//open the file
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//FYI
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		a = string(bs)

		//storing on the server
		f2, err := os.Create(filepath.Join(".", h.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f2.Close()

		_, err = f2.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	<br>`+a)

}
