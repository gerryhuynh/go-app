package download

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

func DownloadConcurrent(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	url, ok := getURLParam(w, query)
	if !ok {
		return
	}

	n, err := strconv.Atoi(query.Get("n"))
	if err != nil {
		http.Error(w, "Invalid or missing n parameter, must be an integer", http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DownloadFile(url)
		}()
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Downloads completed successfully")
}

func FOo() {
	grp := 	&errgroup.Group{}

	for i := 0; i < 10; i++ {
		grp.Go(func() error {
			return nil

		})
}

	grp.Wait()
}
