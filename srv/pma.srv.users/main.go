package main

import (
	"database/sql"
	"fmt"
	mysql "github.com/jfhaines/project_management_app/srv/pma.srv.users/db"
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

	usersDb := mysql.CreateUsersDb(db)
	usersService := service.CreateUsersService(usersDb)

}
