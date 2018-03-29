package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	user       = flag.String("user", "user", "username")
	password   = flag.String("password", "password", "password")
	dbUser     = flag.String("dbUser", "xxx", "username")
	dbPassword = flag.String("dbPassword", "xxx", "password")
	server     = flag.String("server", "localhost", "Server to connect to")
	dbName     = flag.String("dbname", "xxx", "Database name")
	db         *sql.DB
)

func check(u, p string) bool {
	if u == *user && p == *password {
		return true
	}
	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, p, _ := r.BasicAuth()
	if !check(u, p) {
		http.Error(w, "Unauthorised.", 401)
		return
	}
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		for key, values := range r.Form {
			for _, value := range values {
				logPost(key + ":" + value)
			}
		}
		w.Write([]byte("testy test test test"))
	}
}

func logPost(message string) (res sql.Result) {
	_, err := db.Exec("INSERT INTO xxx.dbo.spam(data) VALUES(&message)", sql.Named("message", message))
	if err != nil {
		fmt.Println("From Insert() attempt: " + err.Error())
	}
	return res
}

func main() {
	var err error
	db, err = sql.Open("sqlserver", "server="+*server+";user id="+*dbUser+";password="+*dbPassword)

	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}
	defer db.Close()

	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
