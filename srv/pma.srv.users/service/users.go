package handler

import (
	"context"
	db "github.com/jfhaines/project_management_app/srv/pma.srv.users/db"
	pb "github.com/jfhaines/project_management_app/srv/pma.srv.users/proto"
)

type Service struct {
	usersDb db.UsersDB
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

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.usersDb.GetUser(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{User: user}, nil
}

func (s *Service) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.usersDb.ListUsers()

	if err != nil {
		return nil, err
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := pb.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	createdUser, err := s.usersDb.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: createdUser}, nil
}

func (s *Service) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	user := pb.User{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	editedUser, err := s.usersDb.EditUser(&user)

	if err != nil {
		return nil, err
	}

	return &pb.EditUserResponse{User: editedUser}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.usersDb.DeleteUser(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{}, nil
}
