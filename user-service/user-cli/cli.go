package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/theapemachine/gopilot-api/user-service/proto/user"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "user.json"
)

func parseFile(file string) (*pb.User, error) {
	var user *pb.User

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &user)

	return user, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	user, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	r, err := client.CreateUser(context.Background(), user)

	if err != nil {
		log.Fatalf("could not create: %v", err)
	}

	log.Printf("created: %t", r.Created)

	getAll, err := client.GetUsers(context.Background(), &pb.GetRequest{})

	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}
}
