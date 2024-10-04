package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"go-app/server"
)

func main() {
	const (
		HTTP_PORT = ":8080"
		GRPC_PORT = ":50051"
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("HTTP server starting at http://localhost" + HTTP_PORT)
		if err := http.ListenAndServe(HTTP_PORT, server.HTTPServer()); err != nil {
			log.Printf("HTTP server failed to start: %v\n", err)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("gRPC server starting at http://localhost" + GRPC_PORT)
		if err := server.GRPCServer(GRPC_PORT); err != nil {
			log.Printf("gRPC server failed to serve: %v\n", err)
		}
	}()

	wg.Wait()
}
