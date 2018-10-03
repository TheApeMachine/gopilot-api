package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/theapemachine/gopilot-api/user-service/proto/user"
)

const (
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
	cmd.Init()

	client := pb.NewUserService("go.micro.srv.user", microclient.DefaultClient)
	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	user, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	r, err := client.CreateUser(context.TODO(), user)

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
