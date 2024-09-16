package download

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func DownloadConcurrent(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	n, err := strconv.Atoi(query.Get("n"))
	if err != nil {
		http.Error(w, "Invalid or missing n parameter, must be an integer", http.StatusBadRequest)
		return
	}

	url := query.Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			downloadFile(url)
		}()
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Downloads completed successfully")
}
