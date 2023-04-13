package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jfhaines/project_management_app/database/models"
	users "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DB struct {
	db *sql.DB
}

func CreateUsersDB(db *sql.DB) *DB {
	return &DB{db}
}

type UsersDB interface {
	GetUser(id string) (*users.User, error)
	ListUsers() ([]*users.User, error)
	CreateUser(*users.User) (*users.User, error)
	EditUser(*users.User) (*users.User, error)
	DeleteUser(id string) error
}

func (d *DB) GetUser(id string) (*users.User, error) {
	user, err := models.FindUser(d.db, id)
	fmt.Println("hi")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID %s does not exist", id))
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &users.User{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}, nil
}

func (d *DB) ListUsers() ([]*users.User, error) {
	result, err := models.Users().All(d.db)

	if err != nil {
		return nil, status.Error(codes.Internal, "Unable to retrieve users")
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

func (d *DB) CreateUser(u *users.User) (*users.User, error) {
	newUser := models.User{
		ID:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	err := newUser.Insert(d.db, boil.Infer())
	if err != nil {
		return nil, status.Error(codes.Internal, "Unable to create new user")
	}

	u.Id = newUser.ID

	return u, nil
}

func (d *DB) EditUser(u *users.User) (*users.User, error) {
	user, err := models.FindUser(d.db, u.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID %s does not exist", u.Id))
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	user.Username = u.Username
	user.Password = u.Password
	user.Email = u.Email

	rowsAff, err := user.Update(d.db, boil.Infer())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if rowsAff == 0 {
		return nil, status.Error(codes.Internal, "Failed to update user record")
	}

	return u, nil
}

func (d *DB) DeleteUser(id string) error {
	user, err := models.FindUser(d.db, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return status.Error(codes.NotFound, fmt.Sprintf("User with ID %s does not exist", id))
		} else {
			return status.Error(codes.Internal, err.Error())
		}
	}

	rowsAff, err := user.Delete(d.db)

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	} else if rowsAff == 0 {
		return status.Error(codes.Internal, "Failed to delete user record")
	}

	return nil
}
