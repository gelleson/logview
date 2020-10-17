package search

import (
	"github.com/blevesearch/bleve"
	"time"
)

type bucket struct {
	index   bleve.Index
	created time.Time
}
