package search

import (
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"sync"
	"time"
)

type Option struct {
	TTL time.Duration `json:"ttl"`

	MaxIndicesCapacity int64 `json:"max_indices_capacity"`
	MaxIndicesCount    int64 `json:"max_indices_count"`
}

type searchEngine struct {
	buckets         map[string]*bucket
	option          Option
	maxIndicesCount int64
	mx              sync.Mutex
}

func New(option Option) *searchEngine {
	return &searchEngine{
		buckets: make(map[string]*bucket),
		option:  option,
	}
}

func (search *searchEngine) NewIndex(name string) error {
	search.mx.Lock()
	defer search.mx.Unlock()

	if search.option.MaxIndicesCount != 0 && search.option.MaxIndicesCount < int64(len(search.buckets)) {
		return errors.New(fmt.Sprintf("MaxIndicesCount %d is reached", search.option.MaxIndicesCount))
	}

	mapping := bleve.NewIndexMapping()

	memIndex, err := bleve.NewMemOnly(mapping)

	if err != nil {
		return err
	}

	search.buckets[name] = &bucket{
		created: time.Now(),
		index:   memIndex,
	}

	return nil
}

func (search *searchEngine) DeleteIndex(name string) {
	search.mx.Lock()
	defer search.mx.Unlock()

	delete(search.buckets, name)
}

func (search *searchEngine) Index(name string, entry entry.LogEntry) error {

	index, err := search.getIndex(name)

	if err != nil {
		return err
	}

	count, err := index.index.DocCount()

	if err != nil {
		return err
	}

	if search.option.MaxIndicesCapacity != 0 && search.option.MaxIndicesCapacity <= int64(count) {
		return errors.New(fmt.Sprintf("MaxIndicesCapacity %d is reached", search.option.MaxIndicesCapacity))
	}

	if err := index.index.Index(entry.ID, entry); err != nil {
		return err
	}

	return nil
}

func (search *searchEngine) Search(name string, searchRequest *bleve.SearchRequest) (bleve.SearchResult, error) {

	index, err := search.getIndex(name)

	if err != nil {
		return bleve.SearchResult{}, err
	}

	searchResult, err := index.index.Search(searchRequest)

	if err != nil {
		return bleve.SearchResult{}, err
	}

	return *searchResult, nil
}

func (search *searchEngine) getIndex(key string) (*bucket, error) {
	localIndex, exist := search.buckets[key]

	if !exist {
		return &bucket{}, errors.New(fmt.Sprintf("bucket %s not exist", key))
	}

	return localIndex, nil
}
