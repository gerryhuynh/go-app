package main

import (
	"fmt"
	"net/http"

	"go-app/download"
)

func server() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download)
	return router
}

func main() {
	fmt.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", server()); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
