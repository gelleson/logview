package application

import (
	"github.com/gelleson/logview/pkg/service"
)

type Service struct {
	LogService    *service.LogService
	UploadService *service.UploadService
}

func NewService(logService *service.LogService, uploadService *service.UploadService) *Service {
	return &Service{LogService: logService, UploadService: uploadService}
}
