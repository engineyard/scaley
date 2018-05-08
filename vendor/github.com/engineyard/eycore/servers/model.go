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

package servers

import "fmt"

// Model is an object that represents a server in the Engine Yard Core API.
type Model struct {
	// Standard Server Fields
	ID            int    `json:"id,omitempty"`
	ProvisionedID string `json:"provisioned_id,omitempty"`
	Role          string `json:"role,omitempty"`

	// Server Details
	Dedicated       bool   `json:"dedicated,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	Location        string `json:"location,omitempty"`
	Name            string `json:"name,omitempty"`
	PrivateHostname string `json:"private_hostname,omitempty"`
	PublicHostname  string `json:"public_hostname,omitempty"`
	ReleaseLabel    string `json:"release_label,omitempty"`
	State           string `json:"state,omitempty"`
	EnvironmentURI  string `json:"environment,omitempty"`
	AccountURI      string `json:"account,omitempty"`
	ProviderURI     string `json:"provider,omitempty"`

	// Server flavor is an embedded struct
	Flavor struct {
		ID string `json:"id"`
	} `json:"flavor,omitempty"`

	// Timestamps
	CreatedAt       string `json:"created_at,omitempty"`
	DeletedAt       string `json:"deleted_at,omitempty"`
	DeprovisionedAt string `json:"deprovisioned_at,omitempty"`
	ProvisionedAt   string `json:"provisioned_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`

	// The environment associated with the server. This is tagged for marshaling
	// simply as a convenience and is not actually sent in this form from the
	// API.
	EnvironmentID int `json:"environment_id,omitempty"`
}

// String returns a plain text representation of the Server.
func (model *Model) String() string {
	return fmt.Sprintf("%s (%d) [%s]", model.ProvisionedID, model.ID,
		model.Flavor.ID)
}
