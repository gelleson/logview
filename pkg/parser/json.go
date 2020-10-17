package parser

import (
	"encoding/json"
	"github.com/gelleson/logview/pkg/entry"
	"time"
)

type jsonParser struct {
	level           string
	timestamp       string
	timestampFormat string
}

func NewJSONParser(level string, timestamp string, timestampFormat string) *jsonParser {
	return &jsonParser{level: level, timestamp: timestamp, timestampFormat: timestampFormat}
}

func (p jsonParser) Parse(b []byte) (entry.LogEntry, error) {

	entity := entry.LogEntry{}

	message, err := p.parseJson(b)

	if err != nil {
		return entry.LogEntry{}, err
	}

	timestamp := message[p.timestamp]

	parsedTime, err := p.parseTimestamp(timestamp.(string))

	if err != nil {
		return entry.LogEntry{}, err
	}

	entity.Timestamp = parsedTime

	level, exist := message[p.level]

	if exist {
		entity.Level = level.(string)
	}

	entity.Msg = string(b)

	return entity, nil
}

func (p jsonParser) parseTimestamp(timeRaw string) (time.Time, error) {

	parsedTime, err := time.Parse(p.timestampFormat, timeRaw)

	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func (p jsonParser) parseJson(b []byte) (map[string]interface{}, error) {

	message := make(map[string]interface{})

	err := json.Unmarshal(b, &message)

	if err != nil {
		return map[string]interface{}{}, err
	}

	return message, nil
}
