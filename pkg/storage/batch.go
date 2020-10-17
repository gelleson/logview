package storage

import (
	"github.com/gelleson/logview/pkg/entry"
	"sort"
	"sync"
	"time"
)

type Batch struct {
	records        []entry.LogEntry
	batchSize      int
	totalItems     int
	lastEventTime  time.Time
	lastEvent      entry.LogEntry
	firstEventTime time.Time
	firstEvent     entry.LogEntry
	mx             sync.Mutex
}

func NewBatch() *Batch {
	return &Batch{
		records:   make([]entry.LogEntry, 0),
		batchSize: 250,
	}
}

func (b *Batch) append(item entry.LogEntry) {
	b.mx.Lock()
	defer b.mx.Unlock()

	if b.totalItems == 0 {
		b.firstEvent = item
		b.firstEventTime = item.Timestamp
	}

	if b.totalItems+1 == b.batchSize {
		b.lastEvent = item
		b.lastEventTime = item.Timestamp
	}

	b.records = append(b.records, item)
	b.totalItems++
}

func (b *Batch) isFilled() bool {
	return b.totalItems == 250
}

func (b *Batch) sort() {
	b.mx.Lock()
	defer b.mx.Unlock()

	sort.Slice(b.records, func(i, j int) bool {
		return b.records[i].Timestamp.UnixNano() < b.records[j].Timestamp.UnixNano()
	})

	firstEvent := b.records[0]
	lastEvent := b.records[len(b.records)-1]

	b.firstEventTime = firstEvent.Timestamp
	b.firstEvent = firstEvent
	b.lastEventTime = lastEvent.Timestamp
	b.lastEvent = lastEvent
}

func (b *Batch) BatchSize() int {
	return b.batchSize
}

func (b *Batch) Records() []entry.LogEntry {
	return b.records
}
