package server

import (
	"net/http"

	"go-app/download"
)

func Server() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download) // download one file
	router.HandleFunc("/download-sequential", download.DownloadSequential) // download n files sequentially
	router.HandleFunc("/download-concurrent", download.DownloadConcurrent) // download n files concurrently
	return router
}
