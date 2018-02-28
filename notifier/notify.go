package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/engineyard/scaley/util"
)

func notify(level int, message string) {
	payload := newPayload(level, message)

	if data, err := json.Marshal(payload); err == nil {
		body := bytes.NewReader(data)

		response, postErr := http.Post(util.ReportingURL(), "application/json", body)
		if postErr != nil {
			return
		}

		response.Body.Close()
	}
}
