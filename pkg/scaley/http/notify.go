package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ess/mockable"
)

func notify(url string, level int, message string) {
	payload := newPayload(level, message)

	if mockable.Mocked() {
		fmt.Println(payload)
	} else {
		if data, err := json.Marshal(payload); err == nil {
			body := bytes.NewReader(data)

			response, postErr := http.Post(url, "application/json", body)
			if postErr != nil {
				return
			}

			response.Body.Close()
		}
	}
}
