package main

import (
	"context"
	"log"
	"net"

	pb "github.com/theapemachine/gopilot-api/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.User) (*pb.User, error)
	GetAll() []*pb.User
}

type Repository struct {
	users []*pb.User
}

func (repo *Repository) Create(user *pb.User) (*pb.User, error) {
	updated := append(repo.users, user)
	repo.users = updated
	return user, nil
}

func (repo *Repository) GetAll() []*pb.User {
	return repo.users
}

type service struct {
	repo IRepository
}

func (s *service) CreateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := s.repo.Create(req)

	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, User: user}, nil
}

func (s *service) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	users := s.repo.GetAll()
	return &pb.Response{Users: users}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &service{repo})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
