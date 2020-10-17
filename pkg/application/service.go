package application

import (
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"github.com/gelleson/logview/pkg/services"
	"github.com/gelleson/logview/pkg/storage"
)

type LogService interface {
	SearchString(string, string) (bleve.SearchResult, error)
	ReadLog(string, int) ([]entry.LogEntry, error)
	NewIndex(string, services.Parser, [][]byte) error
	GetLogInfo(string) (storage.IndexInfo, error)
	Append(string, entry.LogEntry) error
}

type Service struct {
	LogService LogService
}

func NewService(logService LogService) *Service {
	return &Service{LogService: logService}
}
