package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jfhaines/project_management_app/database/models"
	users "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (d *DB) GetUser(id string) (*users.User, error) {
	result, err := models.FindUser(d.db, id)

	if err != nil {
		return nil, err
	}

	return &users.User{
		Id:       result.ID,
		Username: result.Username,
		Password: result.Password,
		Email:    result.Email,
	}, nil
}

func (d *DB) ListUsers() ([]*users.User, error) {
	result, err := models.Users().All(d.db)

	if err != nil {
		return nil, err
	}

	var fetchedUsers []*users.User

	for _, u := range result {
		user := &users.User{
			Id:       u.ID,
			Username: u.Username,
			Password: u.Password,
			Email:    u.Email,
		}
		fetchedUsers = append(fetchedUsers, user)
	}

	return fetchedUsers, nil
}

func (d *db) CreateUser(u *users.User) error {
	newUser := models.User{
		ID:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	err := newUser.Insert(d.db, boil.Infer())

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *db) EditUser(u *users.User) error {
	user, err := models.FindUser(d.db, u.Id)

	if err != nil {
		return err
	}

	user.Username = u.Username
	user.Password = u.Password
	user.Email = u.Email

	rowsAffected, err := user.Update(d.db, boil.Infer())

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("User with ID %s not found, no update occurred.", u.Id))
	} else if err != nil {
		return err
	} else {
		return nil
	}
}
