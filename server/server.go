package server

import (
	"net/http"

	"go-app/pkg/download"
	"go-app/pkg/user"
)

func Server() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download)
	router.HandleFunc("/create-user", user.CreateUser)
	return router
}
