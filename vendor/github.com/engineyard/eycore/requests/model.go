// Copyright © 2017 Engine Yard, Inc.
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

package requests

import "fmt"

// Model is an object that represents a background request in the Engine Yard
// Core API.
type Model struct {
	// Standard Request Fields
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Successful bool   `json:"successful,omitempty"`

	// Request Details
	Message       string `json:"message,omitempty"`
	RequestStatus string `json:"request_status,omitempty"`

	// Timestamps
	CreatedAt  string `json:"created_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
	FinishedAt string `json:"finished_at,omitempty"`
	StartedAt  string `json:"started_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`

	EnvironmentID int `json:"environment_id,omitempty"`
}

// String returns a plain text representation of the Request
func (model *Model) String() string {
	status := "In Progress"

	if len(model.FinishedAt) > 0 {
		if model.Successful {
			status = "Successful"
		} else {
			status = "Failed"
		}
	}

	return fmt.Sprintf("%s (%s) <%s> <Finished At %s>", model.Type, model.ID, status, model.FinishedAt)
}
