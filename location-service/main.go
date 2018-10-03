package main

import (
	"context"
	"errors"
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/theapemachine/gopilot-api/location-service/proto/location"
)

type Repository interface {
	FindAll(*pb.Specification) (*pb.Location, error)
}

type LocationRepository struct {
	locations []*pb.Location
}

func (repo *LocationRepository) FindAll(spec *pb.Specification) (*pb.Location, error) {
	for _, location := range repo.locations {
		if spec.Name == location.Name {
			return location, nil
		}
	}

	return nil, errors.New("no location found by that name")
}

type service struct {
	repo Repository
}

func (s *service) FindAll(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	location, err := s.repo.FindAll(req)

	if err != nil {
		return err
	}

	res.Location = location

	return nil
}

func main() {
	locations := []*pb.Location{
		&pb.Location{Name: "Test001"},
	}

	repo := &LocationRepository{locations}

	srv := micro.NewService(
		micro.Name("go.micro.srv.location"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterLocationServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
