package storage

import (
	"errors"
	"github.com/gelleson/logview/pkg/entry"
	"sync"
	"time"
)

type IndexInfo struct {
	Name         string         `json:"name"`
	TotalBatches int            `json:"total_batches"`
	LastEvent    entry.LogEntry `json:"last_event"`
}

type Storage struct {
	logs map[string]*index
	mx   sync.RWMutex
}

func New() *Storage {
	return &Storage{
		logs: make(map[string]*index),
	}
}

func (s *Storage) NewIndex(name string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.logs[name] = newIndex()
}

func (s *Storage) Append(name string, record entry.LogEntry) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	logStorage, exist := s.logs[name]

	if !exist {
		return errors.New("log not found")
	}

	record.Line = logStorage.lastEvent.Line + 1

	logStorage.append(record)
	logStorage.lastEvent = record
	logStorage.sort()

	return nil
}

func (s *Storage) LastBatches(logName string, count int) ([]*Batch, error) {
	logStorage, exist := s.logs[logName]

	if !exist {
		return nil, errors.New("log not found")
	}

	batches := logStorage.last(count)

	return batches, nil
}

func (s *Storage) GetBatches(logName string, skip, total int) ([]*Batch, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return nil, errors.New("log not found")
	}

	batches := logStorage.limit(skip, total)

	return batches, nil
}

func (s *Storage) GetBatch(logName string, skip int) (*Batch, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return nil, errors.New("log not found")
	}

	batches := logStorage.limit(skip, skip+1)

	return batches[0], nil
}

func (s *Storage) GetLastBatches(logName string) ([]*Batch, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return nil, errors.New("log not found")
	}

	batches := logStorage.limit(logStorage.totalBatches-1, logStorage.totalBatches)

	return batches, nil
}

func (s *Storage) FilterByDate(logName string, since, until time.Time) ([]*Batch, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return nil, errors.New("log not found")
	}

	batches := logStorage.filterByDate(since, until)

	return batches, nil
}

func (s *Storage) GetLogInfo(logName string) (IndexInfo, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return IndexInfo{}, errors.New("log not found")
	}

	return IndexInfo{
		Name:         logName,
		TotalBatches: logStorage.totalBatches,
		LastEvent:    logStorage.lastEvent,
	}, nil
}

func (s *Storage) GetLastLineNumber(logName string) (int64, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	logStorage, exist := s.logs[logName]

	if !exist {
		return 0, errors.New("log not found")
	}

	line := logStorage.batches[logStorage.lastBatchIndex-1].lastEvent.Line

	return line, nil
}
