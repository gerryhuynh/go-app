package server

import (
	"net/http"

	"go-app/pkg/download"
	"go-app/pkg/user"
)

func Server() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download) // download one file
	router.HandleFunc("/download-sequential", download.DownloadSequential) // download n files sequentially
	router.HandleFunc("/download-concurrent", download.DownloadConcurrent) // download n files concurrently
	router.HandleFunc("/create-user", user.CreateUser) // create a user
	return router
}
