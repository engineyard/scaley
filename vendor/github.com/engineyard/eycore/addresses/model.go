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

package addresses

import "fmt"

// Model is an object that represents an address retrieved from the
// Engine Yard Core API.
type Model struct {
	// Standard address fields
	ID            int    `json:"id,omitempty"`
	ProvisionedID string `json:"provisioned_id,omitempty"`
	IPAddress     string `json:"ip_address,omitempty"`

	// Address Details
	Server   string `json:"server,omitempty"`
	Location string `json:"location,omitempty"`
	Provider string `json:"provider,omitempty"`

	// Timestamps
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	// The Account associated with the address. This is tagged for marshaling
	// simply as a convenience and is not actually sent in this form from the
	// API.
	AccountID string `json:"account_id,omitempty"`
}

// String returns a plain text representation of the Address.
func (model *Model) String() string {
	return fmt.Sprintf("%s (%s)", model.IPAddress, model.Location)
}
