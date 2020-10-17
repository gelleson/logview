package search

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
)

func NewQueryString(query string) *bleve.SearchRequest {

	stringQuery := bleve.NewQueryStringQuery(query)
	searchRequest := NewSearchRequest(stringQuery)

	return searchRequest
}

func NewSearchRequest(query query.Query) *bleve.SearchRequest {

	searchRequest := bleve.NewSearchRequest(query)

	searchRequest.Highlight = bleve.NewHighlight()
	searchRequest.Size = 1000
	searchRequest.Fields = []string{"*"}

	return searchRequest
}
