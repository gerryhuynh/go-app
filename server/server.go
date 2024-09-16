package server

import (
	"net/http"

	"go-app/download"
)

func Server() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/download", download.Download)
	router.HandleFunc("/download-serial", download.DownloadSerial)
	router.HandleFunc("/download-concurrent", download.DownloadConcurrent)
	return router
}
