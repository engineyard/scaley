package http

import (
	"github.com/engineyard/scaley/pkg/scaley"
)

type LogService struct {
	reportingURL string
}

func NewLogService(reportingURL string) *LogService {
	return &LogService{reportingURL}
}

func (service *LogService) Info(group *scaley.Group, message string) {}

func (service *LogService) Success(group *scaley.Group, message string) {}

func (service *LogService) Failure(group *scaley.Group, message string) {}
