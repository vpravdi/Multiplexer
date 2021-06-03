package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	//http.HandleFunc("/", a)
	http.Handle("/", http.HandlerFunc(a))
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
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error parsing template", err)
	}
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Praveen")
	if err != nil {
		log.Fatal("Error executing template", err)
	}
}
