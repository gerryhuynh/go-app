package download

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	getter "github.com/hashicorp/go-getter/v2"
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
	ctx := context.Background()
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	fmt.Printf("URL: %s\n", u)
	resp, err := http.Head(u)
	if err != nil {
		return fmt.Errorf("failed to fetch headers: %w", err)
	}
	resp.Body.Close()

	fileName := "downloaded_file.zip"
	// Print all headers
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if fn, ok := params["filename"]; ok {
				fileName = fn
			}
		}
	}

	dst := filepath.Join(pwd + "/filedownloads", fileName)

	client := &getter.Client{
		Decompressors: map[string]getter.Decompressor{
			"zip": &getter.ZipDecompressor{},
		},
		Getters: []getter.Getter{
			&getter.HttpGetter{},
		},
	}

	_, err = client.Get(ctx, &getter.Request{
		Src:     u,
		Dst:     dst,
		GetMode: getter.ModeFile,
	})
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	return nil
}
