package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type trial int

func (m trial) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Executing</h1>")

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d trial
	http.ListenAndServe(":8080", d)
}
