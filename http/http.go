package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	user       = flag.String("user", "user", "username")
	password   = flag.String("password", "password", "password")
	dbUser     = flag.String("dbUser", "validate", "username")
	dbPassword = flag.String("dbPassword", "validate", "password")
	server     = flag.String("server", "localhost", "Server to connect to")
	dbName     = flag.String("dbname", "validate", "Database name")
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
				fmt.Fprint(w, key, ",", value)
			}
		}
	}
}

func main() {

	db, err := sql.Open("sqlserver", "server="+*server+";user id="+*dbUser+";password="+*dbPassword)

	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}
	defer db.Close()

	fmt.Println("connected")

	var rows, _ = db.Query("select data from validate.dbo.spam")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	}

	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
