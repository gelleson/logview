package memory

import "time"

type Line struct {
	ID        int    `db:"rowId" json:"id"`
	Line      string `json:"line"`
	Level     string `json:"level"`
	LogName   string `db:"logName" json:"logName"`
	Msg       string `json:"msg"`
	Timestamp string `json:"timestamp"`
}

func NewInfoLine(logName, msg string) Line {
	return Line{
		LogName:   logName,
		Level:     "info",
		Msg:       msg,
		Timestamp: time.Now().String(),
	}
}
