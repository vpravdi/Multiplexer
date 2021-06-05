package main

import (
	"fmt"
	//"github.com/google/uuid"
	"github.com/satori/go.uuid"
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
	//r := mux.NewRouter()
	//http.Handle("/",r)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		//id := uuid.Must(uuid.NewRandom())
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}
