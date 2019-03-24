package document

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `1234567890`)
	}))
	defer ts.Close()

	fetcher := &DocumentFetcher{
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	doc, err := fetcher.fetch(ts.URL)

	assert.Nil(t, err)

	expected := &Document{
		Url:    ts.URL,
		Length: 10,
	}

	assert.Equal(t, expected, doc)
}
