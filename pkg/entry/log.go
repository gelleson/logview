package entry

import "time"

type LogEntry struct {
	ID        string    `json:"id"`
	Line      int64     `json:"line"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Msg       string    `json:"full_line"`
}
