package main

import (
	"context"
	"go-app/pkg/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ExampleServer struct {
	user.UnimplementedUserServiceServer
}

func (s *ExampleServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.Person, error) {
	return &user.Person{
		Id:             "1",
		Email:          "test@example.com",
		Username:       "johndoe",
		FirstName:      "John",
		LastName:       "Doe",
		Age:            30,
		PhoneNumber:    "+1234567890",
		Address:        "123 Main St",
		City:           "New York",
		Country:        "USA",
		PostalCode:     "10001",
		CreatedAt:      1648656000, // Unix timestamp for a sample date
		LastLoginAt:    1648742400, // Unix timestamp for a sample date
		IsActive:       true,
		ProfilePicture: "https://example.com/profile.jpg",
		Occupation:     "Software Engineer",
		Company:        "Tech Corp",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &ExampleServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
