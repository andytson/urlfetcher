package document

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalizerOnSuccessfulDocument(t *testing.T) {

	totals, err := DocumentTotalizer(func(wg *sync.WaitGroup, docs chan Document, errs chan error) {
		docs <- Document{
			Url:    "http://example.org",
			Length: 10,
		}
	})

	assert.Nil(t, err)

	expected := &DocumentTotals{
		TotalBytes:        10,
		numberOfDocuments: 1,
		MiniumSize:        10,
		MaximumSize:       10,
	}

	assert.Equal(t, expected, totals)
}

func TestTotalizerOnMultipleSuccessfulDocument(t *testing.T) {

	totals, err := DocumentTotalizer(func(wg *sync.WaitGroup, docs chan Document, errs chan error) {
		docs <- Document{
			Url:    "http://example.org",
			Length: 10,
		}
		docs <- Document{
			Url:    "http://example.org",
			Length: 20,
		}
	})

	assert.Nil(t, err)

	expected := &DocumentTotals{
		TotalBytes:        30,
		numberOfDocuments: 2,
		MiniumSize:        10,
		MaximumSize:       20,
	}

	assert.Equal(t, expected, totals)
}

func TestTotalizerWithFailedDocument(t *testing.T) {

	totals, err := DocumentTotalizer(func(wg *sync.WaitGroup, docs chan Document, errs chan error) {
		docs <- Document{
			Url:    "http://example.org",
			Length: 10,
		}
		errs <- errors.New("Document failed to download")
	})

	assert.Equal(t, err, errors.New("Document failed to download"))

	assert.Nil(t, totals)
}
