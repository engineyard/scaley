package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ess/mockable"

	"github.com/engineyard/scaley/pkg/util"
)

func notify(level int, message string) {
	if mockable.Mocked() {
		fmt.Println(severity(level), ":", message)
	} else {
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
}
