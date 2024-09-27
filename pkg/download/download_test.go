package download

import (
	"net/url"
	"testing"
)

func TestGetURLParam(t *testing.T) {
	t.Run("Valid URL", func(t *testing.T) {
		want := "https://example.com"
		query := url.Values{"url": []string{want}}
		got, err := getURLParam(query)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("Expected %s, got %s", want, got)
		}
	})

	t.Run("Missing URL", func(t *testing.T) {
		query := url.Values{}
		_, err := getURLParam(query)

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
