package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/theapemachine/gopilot-api/user-service/proto/user"
)

type Repository interface {
	Create(*pb.User) (*pb.User, error)
	GetAll() []*pb.User
}

type UserRepository struct {
	users []*pb.User
}

func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	updated := append(repo.users, user)
	repo.users = updated

	return user, nil
}

func (repo *UserRepository) GetAll() []*pb.User {
	return repo.users
}

type service struct {
	repo Repository
}

func (s *service) CreateUser(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	res.User = user

	return nil
}

func (s *service) GetUsers(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	users := s.repo.GetAll()
	res.Users = users

	return nil
}

func main() {
	repo := &UserRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
