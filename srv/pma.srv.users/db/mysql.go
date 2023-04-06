package mysql

import (
	"database/sql"
	users "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
)

type DB struct {
	db sql.DB
}

func CreateUsersDB(db sql.DB) DB {
	return DB{db}
}

type UsersDB interface {
	GetUser(id string) (*users.User, error)
	ListUsers() (*[]users.User, error)
	CreateUser(*users.User) error
	EditUser(*users.User) error
	DeleteUser(id string) error
}
