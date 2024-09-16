package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	if err := downloadFile(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File downloaded successfully")
}

func downloadFile(u string) error {
	fileName := "foo.zip"

	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	dir, err := getCurrentDirectory()
	if err != nil {
		return err
	}
	dir = dir + "/filedownloads"

	fileName, err = getUniqueFileName(dir, "foo.zip")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func getUniqueFileName(dir, fileName string) (string, error) {
	for i := 1; ; i++ {
		filePath := filepath.Join(dir, fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return fileName, nil
		} else if err != nil {
			return "", err
		}
		fileName = fmt.Sprintf("foo_%d.zip", i)
	}
}

func getCurrentDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
