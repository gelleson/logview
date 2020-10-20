package service

import (
	"bufio"
	"github.com/gelleson/logview/pkg/memory"
	"os"
	"sync"
)

type Inserter interface {
	Insert(...memory.Line) error
}

type UploadService struct {
	inserter Inserter
}

func NewUpload(inserter Inserter) *UploadService {
	return &UploadService{inserter: inserter}
}

func (u *UploadService) UploadFile(fileName string) error {

	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		return err
	}

	batchSize := 1000

	scanner := bufio.NewScanner(file)

	arr := make([]memory.Line, 0)

	for scanner.Scan() {
		arr = append(arr, memory.NewInfoLine(fileName, scanner.Text()))
	}

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i += batchSize {
		i := i
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			size := batchSize + index

			if size > len(arr) {
				size = len(arr) - 1
			}

			_ = u.inserter.Insert(arr[index:size]...)

		}(i)
	}

	wg.Wait()

	return nil
}
