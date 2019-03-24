package document

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type Document struct {
	Url    string
	Length int
}

type DocumentFetcher struct {
	HttpClient *http.Client
}

func (d DocumentFetcher) fetch(url string) (doc *Document, err error) {
	response, err := d.HttpClient.Get(url)

	if err != nil {
		return nil, err
	}
	// The body needs to be read to get its length, as chunked encoding doesn't use Content-Length http header
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	doc = &Document{
		Url:    url,
		Length: len(body),
	}
	return
}

func Fetch(fetcher *DocumentFetcher, url string, wg *sync.WaitGroup, docs chan Document, errs chan error) {
	defer wg.Done()

	doc, err := fetcher.fetch(url)

	if err != nil {
		errs <- err
		return
	}
	docs <- *doc
}
