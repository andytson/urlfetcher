package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/andytson/urlfetcher/internal/document"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Expected 1 argument with the path of the file")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fetcher := &document.DocumentFetcher{
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	totals, err := document.DocumentTotalizer(func(wg *sync.WaitGroup, docs chan document.Document, errs chan error) {
		for scanner.Scan() {
			wg.Add(1)
			go document.Fetch(fetcher, scanner.Text(), wg, docs, errs)
		}

		if err := scanner.Err(); err != nil {
			errs <- err
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Maximum content length: %d\n", totals.MaximumSize)
	fmt.Printf("Minimum content length: %d\n", totals.MiniumSize)
	fmt.Printf("Total content length: %d\n", totals.TotalBytes)
}
