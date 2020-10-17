package application

import (
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"github.com/gelleson/logview/pkg/search"
	"github.com/gelleson/logview/pkg/services"
	"github.com/gelleson/logview/pkg/storage"
	"github.com/labstack/echo"
	"github.com/wailsapp/wails"
)

type Searcher interface {
	NewIndex(string) error
	Index(string, entry.LogEntry) error
	Search(string, *bleve.SearchRequest) (bleve.SearchResult, error)
	DeleteIndex(string)
}

type ControllerBuilder interface {
	Build(*echo.Echo) error
}

type Option struct {
	SearchEngine search.Option `json:"search_engine"`
}

type application struct {
	server   *echo.Echo
	runtime  *wails.Runtime
	searcher Searcher
	storage  *storage.Storage
	service  Service
}

func (a *application) Service() Service {
	return a.service
}

func New(option Option) *application {
	return &application{
		searcher: search.New(option.SearchEngine),
		storage:  storage.New(),
	}
}

func (a *application) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime

	runtime.Dialog.SelectFile()

	return nil
}

func (a *application) Build() error {
	searchService := services.NewLogService(a.searcher, a.storage)

	a.service = Service{
		LogService: searchService,
	}

	return nil
}
