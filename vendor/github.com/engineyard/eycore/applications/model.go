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

package applications

import "fmt"

// Model is an object that represents a application in the Engine Yard Core API.
type Model struct {
	// Standard Application Fields
	ID int `json:"id,omitempty"`

	// Application Details
	Language   string `json:"language,omitempty"`
	Name       string `json:"name,omitempty"`
	Repository string `json:"repository,omitempty"`
	Type       string `json:"type,omitempty"`

	// Timestamps
	CreatedAt string `json:"created_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	// The environment associated with the app. This is tagged for marshaling
	// simply as a convenience and is not actually sent in this form from the
	// API.
	AccountID     string `json:"account_id,omitempty"`
	EnvironmentID int    `json:"environment_id,omitempty"`
}

// String returns a plain text representation of the Application.
func (model *Model) String() string {
	return fmt.Sprintf(
		"%s (%s) [%s]",
		model.Name,
		model.Language,
		model.Repository,
	)
}
