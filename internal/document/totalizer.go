package document

import "sync"

type DocumentTotals struct {
	TotalBytes        int
	numberOfDocuments int
	MiniumSize        int
	MaximumSize       int
}

func DocumentTotalizer(docProducer func(*sync.WaitGroup, chan Document, chan error)) (*DocumentTotals, error) {
	finish := make(chan int)
	docs := make(chan Document)
	errs := make(chan error)
	var wg sync.WaitGroup
	totals := new(DocumentTotals)

	go func() {
		docProducer(&wg, docs, errs)

		wg.Wait()
		close(finish)
	}()

	for {
		select {
		case doc := <-docs:
			totals.TotalBytes += doc.Length
			if totals.numberOfDocuments == 0 || totals.MiniumSize > doc.Length {
				totals.MiniumSize = doc.Length
			}
			if totals.numberOfDocuments == 0 || totals.MaximumSize < doc.Length {
				totals.MaximumSize = doc.Length
			}
			totals.numberOfDocuments++
		case err := <-errs:
			return nil, err
		case <-finish:
			return totals, nil
		}
	}
}
