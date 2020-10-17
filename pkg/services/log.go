package services

import (
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"github.com/gelleson/logview/pkg/search"
	"github.com/gelleson/logview/pkg/storage"
	"github.com/google/uuid"
	"time"
)

type Searcher interface {
	Search(string, *bleve.SearchRequest) (bleve.SearchResult, error)
	NewIndex(string) error
	Index(string, entry.LogEntry) error
}

type Parser interface {
	Parse([]byte) (entry.LogEntry, error)
}

type Persist interface {
	Append(string, entry.LogEntry) error
	NewIndex(string)
	GetLastBatches(logName string) ([]*storage.Batch, error)
	GetBatches(logName string, skip, total int) ([]*storage.Batch, error)
	GetBatch(logName string, skip int) (*storage.Batch, error)
	GetLogInfo(logName string) (storage.IndexInfo, error)
	GetLastLineNumber(logName string) (int64, error)
}

type logService struct {
	searchEngine Searcher
	persist      Persist
}

func NewLogService(searchEngine Searcher, appender Persist) *logService {
	return &logService{
		searchEngine: searchEngine,
		persist:      appender,
	}
}

func (s *logService) SearchString(logName, query string) (bleve.SearchResult, error) {
	return s.searchEngine.Search(logName, search.NewQueryString(query))
}

func (s *logService) ReadLog(name string, batchIndex int) ([]entry.LogEntry, error) {

	result, err := s.persist.GetBatch(name, batchIndex)

	if err != nil {
		return nil, err
	}

	return result.Records(), nil
}

func (s *logService) Append(logName string, record entry.LogEntry) error {
	return s.persist.Append(logName, record)
}

func (s logService) GetLogInfo(logName string) (storage.IndexInfo, error) {
	return s.persist.GetLogInfo(logName)
}

func (s *logService) NewIndex(name string, parser Parser, indices [][]byte) error {
	if err := s.searchEngine.NewIndex(name); err != nil {
		return err
	}

	s.persist.NewIndex(name)

	var lastTime time.Time

	for line, index := range indices {

		logEntry := entry.LogEntry{}

		logEntry, err := parser.Parse(index)

		if err != nil {
			logEntry.Msg = string(index)
			logEntry.Timestamp = lastTime
		}

		if !logEntry.Timestamp.IsZero() {
			lastTime = logEntry.Timestamp
		}

		logEntry.Line = int64(line)
		logEntry.ID = uuid.New().String()

		if err = s.persist.Append(name, logEntry); err != nil {
			return err
		}
		if err = s.searchEngine.Index(name, logEntry); err != nil {
			return err
		}
	}

	return nil
}

func (s *logService) Index(logName string, entry entry.LogEntry) error {
	return s.searchEngine.Index(logName, entry)
}
