package main

import (
	"fmt"
	"log"
	"net/http"

	"go-app/server"

	"golang.org/x/sync/errgroup"
)

func main() {
	const (
		HTTP_PORT = ":8080"
		GRPC_PORT = ":50051"
	)

	eg := &errgroup.Group{}

	eg.Go(func() error {
		fmt.Println("HTTP server starting at port", HTTP_PORT)
		return fmt.Errorf("HTTP: %w", http.ListenAndServe(HTTP_PORT, server.HTTPServer()))
	})

	eg.Go(func() error {
		fmt.Println("gRPC server starting at port", GRPC_PORT)
		return fmt.Errorf("gRPC: %w", server.GRPCServer(GRPC_PORT))
	})

	if err := eg.Wait(); err != nil {
		log.Printf("Server failed to start: %v\n", err)
	}
}
