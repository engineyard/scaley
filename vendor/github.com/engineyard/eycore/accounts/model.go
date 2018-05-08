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

package accounts

import "fmt"

// Model is an object that represents an account retrieved from the
// Engine Yard Core API.
type Model struct {
	// Standard account fields
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Account Details
	Plan        string `json:"plan,omitempty"`
	SupportPlan string `json:"support_plan,omitempty"`
	Type        string `json:"type,omitempty"`

	// Timestamps
	CanceledAt  string `json:"canceled_at,omitempty"`
	CancelledAt string `json:"cancelled_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`

	// The User associated with the account. This is tagged for marshaling
	// simply as a convenience and is not actually sent in this form from the
	// API.
	UserID string `json:"user_id,omitempty"`
}

// String returns a plain text representation of the Model.
func (model *Model) String() string {
	return fmt.Sprintf("%s (%s)", model.Name, model.ID)
}