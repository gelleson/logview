package service

import (
	"github.com/gelleson/logview/pkg/memory"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

type SuiteUploadService struct {
	suite.Suite
	upload *UploadService
	memory *memory.Memory
}

func (s *SuiteUploadService) SetupTest() {

	s.memory, _ = memory.New()

	s.upload = NewUpload(s.memory)
}

func (s SuiteUploadService) TestUploadFile() {
	file, err := ioutil.TempFile("", "access.*.log")

	s.Assert().Nil(err)

	defer os.Remove(file.Name())

	for i := 0; i < 1000; i++ {
		file.Write([]byte("2020-10-11T08:48:00.117Z 2020-10-11 08:48:00.117  INFO 1 --- [   scheduling-1] ---------------- ted\n"))
	}

	err = s.upload.UploadFile(file.Name())

	s.Assert().Nil(err)
}

func TestRunUploadSuite(t *testing.T) {
	suite.Run(t, new(SuiteUploadService))
}

func BenchmarkUpload_UploadFile(b *testing.B) {
	m, _ := memory.New()

	upload := NewUpload(m)
	file, _ := ioutil.TempFile("", "access.*.log")

	defer os.Remove(file.Name())

	for i := 0; i < 1000; i++ {
		file.Write([]byte("2020-10-11T08:48:00.117Z 2020-10-11 08:48:00.117  INFO 1 --- [   scheduling-1] ---------------- ted\n"))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		upload.UploadFile(file.Name())
		m, _ := memory.New()

		upload = NewUpload(m)
	}
}
