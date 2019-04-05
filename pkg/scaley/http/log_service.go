package http

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/scaley"
)

type LogService struct {
	reportingURL string
}

func NewLogService(reportingURL string) *LogService {
	return &LogService{reportingURL}
}

func (service *LogService) Info(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		warning,
		normalize(group, message),
	)
}

func (service *LogService) Success(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		okay,
		normalize(group, message),
	)
}

func (service *LogService) Failure(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		failure,
		normalize(group, message),
	)
}

func normalize(group *scaley.Group, message string) string {
	return fmt.Sprintf("Group[%s]: %s", group.Name, message)
}
