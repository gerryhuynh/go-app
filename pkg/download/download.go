package download

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const downloadDir = "filedownloads"
const defaultFileName = "foo.zip"

func Download(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	url, err := getURLParam(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	n, err := getNParam(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s, err := getSequentialParam(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if s {
		downloadSequential(w, url, n)
	} else {
		downloadConcurrent(w, r, url, n)
	}
}

func downloadSequential(w http.ResponseWriter, url string, n int) {
	resp, err := downloadFromURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	buffer, err := createBuffer(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dir, err := getDownloadDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < n; i++ {
		file, err := createFile(dir, defaultFileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, bytes.NewReader(buffer.Bytes()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File downloaded successfully")
}

func downloadConcurrent(w http.ResponseWriter, r *http.Request, url string, n int) {
	panic("not yet implemented")
}

func createBuffer(resp *http.Response) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func createFile(dir, fileName string) (*os.File, error) {
	filePath, err := getUniqueFilePath(dir, fileName)
	if err != nil {
		return nil, err
	}

	return os.Create(filePath)
}

func getURLParam(query url.Values) (string, error) {
	url := query.Get("url")
	if url == "" {
		return "", fmt.Errorf("URL parameter is required")
	}
	return url, nil
}

func getNParam(query url.Values) (int, error) {
	nStr := query.Get("n")
	if nStr == "" {
		return 1, nil
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, fmt.Errorf("invalid n parameter: %w", err)
	}
	return n, nil
}

func getSequentialParam(query url.Values) (bool, error) {
	sStr := query.Get("s")
	if sStr == "" {
		return false, nil
	}

	s, err := strconv.ParseBool(sStr)
	if err != nil {
		return false, fmt.Errorf("invalid sequential parameter: %w", err)
	}
	return s, nil
}

func downloadFromURL(u string) (*http.Response, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	return resp, nil
}

func getDownloadDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir, downloadDir)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return dir, nil
}

func getUniqueFilePath(dir, fileName string) (string, error) {
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	ext := filepath.Ext(fileName)

	for i := 1; ; i++ {
		filePath := filepath.Join(dir, fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return filePath, nil
		} else if err != nil {
			return "", err
		}
		fileName = fmt.Sprintf("%s_%d%s", baseName, i, ext)
	}
}
