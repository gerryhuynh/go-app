package main

import (
	"context"
	"fmt"
	"go-app/pkg/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	user, err := client.GetUser(context.Background(), &user.GetUserRequest{})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	userBytes, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	fmt.Println("User:", user)
	fmt.Println("\nUser Bytes:", userBytes)
	fmt.Println("\nLength:", len(userBytes))
	fmt.Println("\nString:", string(userBytes))
}
