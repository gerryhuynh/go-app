package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-app/pkg/download"
)

func TestServer(t *testing.T) {
	server := Server()
	restoreDownloadFile := mockDownloadFile()
	defer restoreDownloadFile()

	t.Run("Server returns non-nil handler", func(t *testing.T) {
		if server == nil {
			t.Error("Server() returned nil, want non-nil")
		}
	})

	t.Run("Server handles /download route", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/download?url=https://example.com/file.zip", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Server returned %d for /download route, want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Server handles /download-sequential route", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/download-sequential?url=https://example.com/file.zip&n=3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Server returned %d for /download-sequential route, want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Server handles /download-concurrent route", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/download-concurrent?url=https://example.com/file.zip&n=3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Server returned %d for /download-concurrent route, want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Server handles /create-user route", func(t *testing.T) {
		userData := `{"id": "1", "email": "john@example.com"}`
		request := httptest.NewRequest(http.MethodPost, "/create-user", strings.NewReader(userData))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
				t.Errorf("Server returned %d for /create-user route, want %d", response.Code, http.StatusCreated)
		}
	})

	t.Run("Server returns 404 for unknown route", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/unknown", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("Server returned %d for unknown route, want %d", response.Code, http.StatusNotFound)
		}
	})
}

func mockDownloadFile() func() {
	originalDownloadFile := download.DownloadFile

	download.DownloadFile = func(u string) error {
		return nil
	}

	return func() {
		download.DownloadFile = originalDownloadFile
	}
}
