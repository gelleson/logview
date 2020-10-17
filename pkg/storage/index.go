package storage

import (
	"github.com/gelleson/logview/pkg/entry"
	"sort"
	"sync"
	"time"
)

const batchSize = 250

type index struct {
	batches        []*Batch
	lastBatchIndex int
	lastEvent      entry.LogEntry
	size           int
	totalBatches   int
	mx             sync.Mutex
}

func newIndex() *index {
	return &index{}
}

func (i *index) append(log entry.LogEntry) {

	if i.lastBatchIndex == 0 {
		i.batches = append(i.batches, NewBatch())
		i.lastBatchIndex++
	}

	localBatch := i.batches[i.lastBatchIndex-1]

	if localBatch.isFilled() {
		localBatch.sort()
		localBatch = NewBatch()
		i.batches = append(i.batches, localBatch)
		i.totalBatches++
		i.lastBatchIndex++
	}

	localBatch.append(log)
	localBatch.lastEvent = log
	localBatch.lastEventTime = log.Timestamp

	if localBatch.isFilled() {
		localBatch.sort()
		localBatch = NewBatch()
		i.batches = append(i.batches, localBatch)
		i.totalBatches++
		i.lastBatchIndex++
	}

	i.lastEvent = log
}

func (i *index) limit(skip, until int) []*Batch {

	return i.batches[skip:until]
}

func (i *index) last(count int) []*Batch {

	batchEndIndex := count / batchSize

	return i.batches[i.totalBatches-count : batchEndIndex]
}

func (i *index) sort() {
	i.mx.Lock()
	defer i.mx.Unlock()

	for _, b := range i.batches {
		b.sort()
	}

	sort.Slice(i.batches, func(index, jindex int) bool {
		return i.batches[index].lastEventTime.UnixNano() < i.batches[jindex].lastEventTime.UnixNano()
	})
}

func (i *index) filterByDate(since, until time.Time) []*Batch {

	i.sort()

	sinceFound := false
	sinceIndex := 0
	lastIndex := len(i.batches) - 1

Loop:
	for index, batch := range i.batches {

		if since.After(batch.lastEventTime) && !sinceFound {
			sinceFound = false
			sinceIndex = index
		}

		if until.Before(batch.lastEventTime) {
			lastIndex = index
			break Loop
		}
	}

	return i.batches[sinceIndex:lastIndex]
}
