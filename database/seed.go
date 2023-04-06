package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		"root",                   // username
		"",                       // pass
		"127.0.0.1",              // host
		"3306",                   // port
		"project_management_app", // db name
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

	files, err := ioutil.ReadDir("./sql")
	for _, f := range files {
		file, err := os.Open("./sql/" + f.Name())
		if err != nil {
			fmt.Println(f.Name(), ": ", err)
		} else {
			fmt.Println("Executing" + " " + f.Name())
			scanner := bufio.NewScanner(file)
			scanner.Split(querySplitFunc)

			for scanner.Scan() {
				db.Exec(scanner.Text())
			}

			file.Close()
		}
	}
}

func querySplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return
	}

	if i := strings.Index(string(data), ";"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
