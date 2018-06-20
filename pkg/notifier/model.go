package notifier

import (
	"fmt"
	"os"
)

const (
	okay int = iota
	warning
	failure
)

type data struct {
	Type     string `json:"Type"`
	Severity string `json:"Severity"`

	CurrentValue string `json:"CurrentValue"`
	FailureMax   string `json:"FailureMax"`

	RawMessage string `json:"raw_message"`
}

type payload struct {
	Message string `json:"message"`

	Data *data `json:"data"`
}

func (p *payload) String() string {
	return fmt.Sprintf("%s : %s", p.Data.Severity, p.Data.RawMessage)
}

func newPayload(level int, message string) *payload {
	return &payload{
		Message: "alert",
		Data: &data{
			Type:         fmt.Sprintf("process-scaley[%d]", os.Getpid()),
			RawMessage:   message,
			FailureMax:   "1.0",
			CurrentValue: currentValue(level),
			Severity:     severity(level),
		},
	}
}

func severity(level int) string {
	if level == failure {
		return "FAILURE"
	}

	if level == warning {
		return "WARNING"
	}

	return "OKAY"
}

func currentValue(level int) string {
	if level == failure {
		return "1.0"
	}

	return "0.0"
}
