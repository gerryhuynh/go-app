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

	"golang.org/x/sync/errgroup"
)

const (
	DOWNLOAD_DIR      = "filedownloads"
	DEFAULT_FILE_NAME = "foo.zip"
)

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
		downloadConcurrent(w, url, n)
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

	dir, err := createNewDownloadDir()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < n; i++ {
		file, err := createFile(dir, DEFAULT_FILE_NAME, i)
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

func downloadConcurrent(w http.ResponseWriter, url string, n int) {
	var buffer *bytes.Buffer
	var dir string

	eg := &errgroup.Group{}

	eg.Go(func() error {
		resp, err := downloadFromURL(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		buffer, err = createBuffer(resp)
		return err
	})

	eg.Go(func() error {
		var err error
		dir, err = createNewDownloadDir()
		return err
	})

	if err := eg.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	eg = &errgroup.Group{}

	for i := 0; i < n; i++ {
		i := i
		eg.Go(func() error {
			file, err := createFile(dir, DEFAULT_FILE_NAME, i)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, bytes.NewReader(buffer.Bytes()))
			return err
		})
	}

	if err := eg.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "File downloaded successfully")
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

func createBuffer(resp *http.Response) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func createNewDownloadDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir, DOWNLOAD_DIR)

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		if err := os.RemoveAll(dir); err != nil {
			return "", fmt.Errorf("failed to remove directory: %w", err)
		}
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return dir, nil
}

func createFile(dir, fileName string, i int) (*os.File, error) {
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	ext := filepath.Ext(fileName)

	filePath := filepath.Join(dir, fmt.Sprintf("%s_%d%s", baseName, i, ext))

	return os.Create(filePath)
}
