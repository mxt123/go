package main

import (
	"fmt"
	"net/http"
)

func check(user, pass string) bool {
	fmt.Printf(user + ":" + pass + "\n")
	if user == "user" && pass == "password" {
		return true
	} else {
		return false

	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	user, pass, _ := r.BasicAuth()
	if !check(user, pass) {
		http.Error(w, "Unauthorised.", 401)
		return
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
