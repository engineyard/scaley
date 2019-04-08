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

// Copyright © 2019 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
