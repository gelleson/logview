package application

import (
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"github.com/gelleson/logview/pkg/memory"
	"github.com/gelleson/logview/pkg/service"

	"github.com/wailsapp/wails"
)

type Searcher interface {
	NewIndex(string) error
	Index(string, entry.LogEntry) error
	Search(string, *bleve.SearchRequest) (bleve.SearchResult, error)
	DeleteIndex(string)
}

type ControllerBuilder interface {
}

type Option struct {
}

type application struct {
	runtime  *wails.Runtime
	searcher Searcher
	memory   *memory.Memory
	service  *Service
}

func (a *application) Service() *Service {
	return a.service
}

func New(option Option) *application {
	m, _ := memory.New()

	return &application{
		memory: m,
	}
}

func (a *application) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime

	return nil
}

func (a *application) Build() error {
	searchService := service.NewLogService(a.memory)

	logReaderService := service.NewUpload(a.memory)

	a.service = NewService(searchService, logReaderService)

	return nil
}
