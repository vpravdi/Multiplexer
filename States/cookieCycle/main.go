package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set"> Set the cookie</a></h1>`)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "Session",
		Value: "ohllala",
	})
	fmt.Fprintln(res, `<h1><a href="/read">Read the cookie</a></h1>`)

}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Session")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(res, `<h1>Your cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Session")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return

	}
	c.MaxAge = -1
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)

}
