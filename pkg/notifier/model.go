// Copyright © 2018 Engine Yard, Inc.
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
