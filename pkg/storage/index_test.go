package storage

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestIndex_append(t *testing.T) {

	indx := index{
		batches: make([]*Batch, 0),
	}

	entries := gen()

	old := genOld()

	for _, entry := range entries {
		indx.append(entry)
	}

	for _, entry := range old {
		indx.append(entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.UnixNano() > entries[j].Timestamp.UnixNano()
	})

	sort.Slice(old, func(i, j int) bool {
		return old[i].Timestamp.UnixNano() > old[j].Timestamp.UnixNano()
	})

	fmt.Println(len(old), entries[0].Timestamp, old[0].Timestamp)
	fmt.Println(len(entries))

	date := indx.filterByDate(time.Date(2032, 1, 1, 1, 1, 1, 1, time.Local), time.Now())

	fmt.Println(date[0])
}

func BenchmarkAppend(b *testing.B) {
	indx := index{
		batches: make([]*Batch, 0),
	}

	entries := gen()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indx.append(entries[0])
	}
}
