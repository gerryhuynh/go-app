package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const downloadDir = "filedownloads"
const defaultFileName = "foo.zip"

var fileMutex sync.Mutex

func Download(w http.ResponseWriter, r *http.Request) {
	url, ok := getURLParam(w, r.URL.Query())
	if !ok {
		return
	}

	if err := DownloadFile(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File downloaded successfully")
}

var DownloadFile = func(u string) error {
	resp, err := downloadFromURL(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dir, err := getDownloadDir()
	if err != nil {
		return err
	}

	fileMutex.Lock()
	defer fileMutex.Unlock()

	fileName, err := getUniqueFileName(dir, defaultFileName)
	if err != nil {
		return err
	}

	if err := saveFile(dir, fileName, resp); err != nil {
		return err
	}

	fmt.Printf("File %q downloaded\n", fileName)
	return nil
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

func getUniqueFileName(dir, fileName string) (string, error) {
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	ext := filepath.Ext(fileName)

	for i := 1; ; i++ {
		filePath := filepath.Join(dir, fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return fileName, nil
		} else if err != nil {
			return "", err
		}
		fileName = fmt.Sprintf("%s_%d%s", baseName, i, ext)
	}
}

func saveFile(dir, fileName string, resp *http.Response) error {
	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
