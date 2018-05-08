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

package users

import "fmt"

// Model is an object that represents a user in the Engine Yard Core API.
type Model struct {
	// Standard User Fields
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`

	// User Details
	Role     string `json:"role,omitempty"`
	APIToken string `json:"api_token,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	Staff    bool   `json:"staff,omitempty"`

	// Timestamps
	CreatedAt string `json:"created_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// String returns a plain text representation of the User
func (model *Model) String() string {
	return fmt.Sprintf("%s <%s>", model.Name, model.Email)
}
