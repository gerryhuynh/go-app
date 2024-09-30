package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"go-app/pkg/download"
	"go-app/pkg/user"

	"google.golang.org/grpc"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.Person, error) {
	return &user.Person{
		Id:             int32(1),
		Email:          "test@test.com",
		Username:       "testuser",
		FirstName:      "John",
		LastName:       "Doe",
		Age:            int32(30),
		PhoneNumber:    "+1234567890",
		Address:        "123 Main St",
		City:           "Anytown",
		Country:        "USA",
		PostalCode:     "12345",
		CreatedAt:      time.Now().Unix(),
		LastLoginAt:    time.Now().Unix(),
		IsActive:       true,
		ProfilePicture: "https://example.com/profile.jpg",
		Occupation:     "Software Developer",
		Company:        "Tech Corp",
	}, nil
}

func HTTPServer() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download)
	router.HandleFunc("/create-user", user.CreateUser)
	router.HandleFunc("/marshal-user", user.MarshalUser)
	return router
}

func GRPCServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &UserServer{})

	return grpcServer.Serve(lis)
}
