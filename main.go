package main

import (
	"fmt"
	"net/http"

	"go-app/server"
)

func main() {
	fmt.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", server.Server()); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
