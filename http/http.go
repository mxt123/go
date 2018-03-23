package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	user     = flag.String("user", "user", "username")
	password = flag.String("password", "password", "password")
)

func check(u, p string) bool {
	if u == *user && p == *password {
		return true
	} else {
		return false

	}
}

//http://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
func handler(w http.ResponseWriter, r *http.Request) {
	u, p, _ := r.BasicAuth()
	if !check(u, p) {
		http.Error(w, "Unauthorised.", 401)
		return
	} else {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, r.URL.Path[1:])
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			for key, values := range r.Form {
				for _, value := range values {
					fmt.Fprint(w, key, ",", value)
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
