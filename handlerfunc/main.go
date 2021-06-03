package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	//http.HandleFunc("/", a)
	http.Handle("/", http.HandlerFunc(a))
	http.HandleFunc("/kaka.jpeg", kaka)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)

}

func a(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "a ran")

}

func kaka(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("kaka.jpeg")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	defer f.Close()

	io.Copy(res, f)
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
