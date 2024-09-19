package download

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadInMemory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	n, err := strconv.Atoi(query.Get("n"))
	if err != nil {
		http.Error(w, "Invalid or missing n parameter, must be an integer", http.StatusBadRequest)
		return
	}

	url, ok := getURLParam(w, query)
	if !ok {
		return
	}

	buffer := bytes.NewBuffer([]byte{})
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	io.Copy(buffer, resp.Body)

	// create n files
	files := make([]*os.File, n)
	for i := 0; i < n; i++ {
		files[i], err = os.Create(fmt.Sprintf("file%d.txt", i))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// io.File -> io.Writer
	writers := make([]io.Writer, n)
	for i := 0; i < n; i++ {
		writers[i] = files[i]
	}

	out := io.MultiWriter(
		writers...,
	)

	io.Copy(out, buffer)
	for _, file := range files {
		file.Close()
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Downloads completed successfully")
}
