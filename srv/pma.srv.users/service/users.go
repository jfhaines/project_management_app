package handler

import (
	"context"
	db "github.com/jfhaines/project_management_app/srv/pma.srv.users/db"
	pb "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
)

type Service struct {
	usersDb db.UsersDb
}

func CreateUsersService(d db.UsersDB) Service {
	return Service{usersDb: d}
}

type UsersService interface {
	GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error)
	ListUsers(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	EditUser(context.Context, *pb.EditUserRequest) (*pb.EditUserResponse, error)
	DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}
