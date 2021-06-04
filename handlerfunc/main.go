package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/b", http.HandlerFunc(a))
	http.HandleFunc("/kaka.jpeg", kaka)
	http.HandleFunc("/me/", myName)
	http.HandleFunc("/c", fileServe)
	http.ListenAndServe(":8080", nil)

}

func a(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "a ran")

}

func fileServe(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "kaka.jpeg")
}

func kaka(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("kaka.jpeg")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)

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
