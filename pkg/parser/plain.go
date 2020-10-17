package parser

import (
	"github.com/btubbs/datetime"
	"github.com/gelleson/logview/pkg/entry"
	"regexp"
)

type plainText struct {
	reg *regexp.Regexp
}

func NewPlainTextParser(query string) (*plainText, error) {
	compiledReg, err := regexp.Compile(query)

	if err != nil {
		return nil, err
	}

	return &plainText{
		reg: compiledReg,
	}, nil
}

func (p plainText) Parse(b []byte) (entry.LogEntry, error) {

	entry := entry.LogEntry{}

	names := p.reg.SubexpNames()

	result := p.reg.FindAllStringSubmatch(string(b), -1)

	m := make(map[string]string)

	for i, n := range result[0] {
		m[names[i]] = n
	}

	timestamp := m["timestamp"]

	local, err := datetime.ParseLocal(timestamp)

	if err == nil {
		entry.Timestamp = local
	}

	level, exist := m["level"]

	if exist {
		entry.Level = level
	}

	entry.Msg = string(b)

	return entry, nil
}
