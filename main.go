package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"go-app/server"
)

func main() {
	useGRPC := flag.Bool("grpc", false, "Use gRPC server")
	flag.Parse()

	if *useGRPC {
		fmt.Println("gRPC server starting at http://localhost:50051")
		if err := server.GRPCServer(); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	} else {
		fmt.Println("HTTP server starting at http://localhost:8080")
		if err := http.ListenAndServe(":8080", server.HTTPServer()); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	}
}
