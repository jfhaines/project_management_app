package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfhaines/project_management_app/srv/pma.srv.users/config"
	mysql "github.com/jfhaines/project_management_app/srv/pma.srv.users/db"
	pb "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
	service "github.com/jfhaines/project_management_app/srv/pma.srv.users/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DB_USERNAME,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error: cannot connect to database\n%v", err))
	}

	usersDb := mysql.CreateUsersDB(db)
	usersService := service.CreateUsersService(usersDb)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.SERVICE_PORT))

	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &usersService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
