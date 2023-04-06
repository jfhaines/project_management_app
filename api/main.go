package main

import (
	"net/http"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		"root",      // username
		"",          // pass
		"127.0.0.1", // host
		"3306",      // port
		"test",      // db name
	))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error: cannot connect to database\n%v", err)
		return
	}

	db.QueryRow("SELECT `id`, `name` FROM cats WHERE id = 2")

	http.ListenAndServe(":3333", nil)
}