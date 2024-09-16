package download

import (
	"net/http"
	"net/url"
)

func getURLParam(w http.ResponseWriter, query url.Values) (string, bool) {
	url := query.Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return url, false
	}
	return url, true
}
