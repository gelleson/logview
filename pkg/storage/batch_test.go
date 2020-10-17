package storage

import (
	"github.com/gelleson/logview/pkg/entry"
	"math/rand"
	"testing"
	"time"
)

func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func randateWithDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2001, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func gen() []entry.LogEntry {

	entries := make([]entry.LogEntry, 255)

	for i := 0; i < 255; i++ {
		entries[i] = entry.LogEntry{
			ID:        "as",
			Line:      int64(i),
			Timestamp: randate(),
			Level:     "",
			Msg:       "",
		}
	}

	return entries
}

func genOld() []entry.LogEntry {

	entries := make([]entry.LogEntry, 255)

	for i := 0; i < 255; i++ {
		entries[i] = entry.LogEntry{
			ID:        "as",
			Line:      int64(i),
			Timestamp: randateWithDate(),
			Level:     "",
			Msg:       "",
		}
	}

	return entries
}

func BenchmarkBench_sort(b *testing.B) {
	entries := gen()

	batch := NewBatch()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		batch.records = entries
		batch.sort()
	}
}
