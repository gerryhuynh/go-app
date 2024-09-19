package main

import (
	"context"
	"fmt"
	"go-app/pkg/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	user, err := client.GetUser(context.Background(), &user.GetUserRequest{})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	userBytes, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("failed to marshal user: %v", err)
	}

	fmt.Printf("User: %v\n", string(userBytes)	)
}
