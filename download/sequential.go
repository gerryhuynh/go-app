package download

import (
	"fmt"
	"net/http"
	"strconv"
)

func DownloadSequential(w http.ResponseWriter, r *http.Request) {
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

	for i := 0; i < n; i++ {
		if err := downloadFile(url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Downloads completed successfully")
}
