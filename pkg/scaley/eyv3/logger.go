package eyv3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ess/mockable"

	"github.com/engineyard/scaley/pkg/scaley"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Info(group scaley.Group, message string) {
	logger.notify(scaley.Warning, logger.normalize(group, message))
}

func (logger *Logger) Success(group scaley.Group, message string) {
	logger.notify(scaley.Okay, logger.normalize(group, message))
}

func (logger *Logger) Failure(group scaley.Group, message string) {
	logger.notify(scaley.Failure, logger.normalize(group, message))
}

func (logger *Logger) normalize(group scaley.Group, message string) string {
	return fmt.Sprintf("Group[%s]: %s", group.Name, message)
}

func (logger *Logger) notify(level scaley.Severity, message string) {
	payload := newLogPayload(level, message)

	if data, err := json.Marshal(payload); err == nil {
		body := bytes.NewReader(data)

		response, postErr := http.Post(util.ReportingURL(), "application/json", body)

		if postErr != nil {
			return
		}

		response.Body.Close()
	}
}

type logData struct {
	Type     string `json:"Type"`
	Severity string `json:"Severity"`

	CurrentValue string `json:"CurrentValue"`
	FailureMax   string `json:"FailureMax"`

	RawMessage string `json:"raw_message"`
}

type logPayload struct {
	Message string `json:"message"`

	Data *logData `json:"data"`
}

func newLogPayload(level scaley.Severity, message string) *logPayload {
	value := "0.0"

	if level == scaley.Failure {
		value = "1.0"
	}

	return &logPayload{
		Message: "alert",
		Data: &logData{
			Type:         fmt.Sprintf("process-scaley[%d]", os.Getpid()),
			RawMessage:   message,
			FailureMax:   "1.0",
			CurrentValue: value,
			Severity:     level.String(),
		},
	}
}
