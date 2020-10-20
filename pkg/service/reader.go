package service

import "github.com/gelleson/logview/pkg/memory"

type LogReader interface {
	Offset(string, int, int) ([]memory.Line, error)
	Match(string, string) ([]memory.Line, error)
	LogsList() ([]memory.LogInfo, error)
}

type LogService struct {
	logReader LogReader
}

func NewLogService(logReader LogReader) *LogService {
	return &LogService{logReader: logReader}
}

func (s *LogService) Match(logName string, query string) ([]memory.Line, error) {
	return s.logReader.Match(logName, query)
}

func (s *LogService) Offset(logName string, skip, offset int) ([]memory.Line, error) {
	return s.logReader.Offset(logName, skip, offset)
}

func (s *LogService) LogsList() ([]memory.LogInfo, error) {
	return s.logReader.LogsList()
}
