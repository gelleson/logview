package controller

import (
	"bufio"
	"github.com/blevesearch/bleve"
	"github.com/gelleson/logview/pkg/entry"
	"github.com/gelleson/logview/pkg/parser"
	"github.com/gelleson/logview/pkg/services"
	"github.com/gelleson/logview/pkg/storage"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"strconv"
)

type LogService interface {
	SearchString(string, string) (bleve.SearchResult, error)
	ReadLog(string, int) ([]entry.LogEntry, error)
	NewIndex(string, services.Parser, [][]byte) error
	GetLogInfo(string) (storage.IndexInfo, error)
	Append(string, entry.LogEntry) error
}

type ReadLogController struct {
	searchService LogService
}

func NewReadLogController(searchService LogService) *ReadLogController {
	return &ReadLogController{searchService: searchService}
}

func (sc *ReadLogController) Build(e *echo.Echo) error {

	group := e.Group("/logs")

	group.POST("/", sc.CreateIndex)
	group.GET("/:logName", sc.LogInfoController)
	group.POST("/:logName", sc.AppendLines)
	group.GET("/:logName/:batchId", sc.GetBatchController)
	group.GET("/:logName/search", sc.SearchController)

	return nil
}

func (sc *ReadLogController) SearchController(c echo.Context) error {

	param := c.Get("logName")

	queryParam := c.QueryParam("query")

	searchResult, err := sc.searchService.SearchString(param.(string), queryParam)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, searchResult)
}

func (sc *ReadLogController) LogInfoController(c echo.Context) error {

	param := c.Param("logName")

	searchResult, err := sc.searchService.GetLogInfo(param)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, searchResult)
}

func (sc *ReadLogController) GetBatchController(c echo.Context) error {

	param := c.Param("logName")
	batchId := c.Param("batchId")

	index, err := strconv.Atoi(batchId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	searchResult, err := sc.searchService.ReadLog(param, index)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, searchResult)
}

func (sc *ReadLogController) CreateIndex(c echo.Context) error {

	file, err := c.FormFile("file")

	if err != nil {
		return err
	}
	src, err := file.Open()

	newBuffer := bufio.NewReader(src)

	fileMsg := make([][]byte, 0)

	for {
		line, _, err := newBuffer.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}

			return echo.NewHTTPError(http.StatusBadGateway, err)
		}

		fileMsg = append(fileMsg, line)
	}

	value := c.FormValue("schema")

	textParser, err := parser.NewPlainTextParser(value)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	if err = sc.searchService.NewIndex(file.Filename, textParser, fileMsg); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "ok",
	})
}

func (sc *ReadLogController) AppendLines(c echo.Context) error {

	param := c.Param("logName")

	entities, err := readFile(c)

	if err != nil {
		return err
	}

	for _, entity := range entities {
		if err := sc.searchService.Append(param, entity); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "ok",
	})
}

func readFile(c echo.Context) ([]entry.LogEntry, error) {
	file, err := c.FormFile("file")

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadGateway, err)
	}

	src, err := file.Open()

	newBuffer := bufio.NewReader(src)

	fileMsg := make([][]byte, 0)

	for {
		line, _, err := newBuffer.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, echo.NewHTTPError(http.StatusBadGateway, err)
		}

		fileMsg = append(fileMsg, line)
	}

	value := c.FormValue("schema")

	textParser, err := parser.NewPlainTextParser(value)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadGateway, err)
	}

	messages := make([]entry.LogEntry, 0)

	for _, payload := range fileMsg {

		message, err := textParser.Parse(payload)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}
