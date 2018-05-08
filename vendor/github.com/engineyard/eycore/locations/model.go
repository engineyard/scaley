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

package locations

import "fmt"

// Limits is an object that describes the upper server and address limits for
// the location that embeds the object.
type Limits struct {
	Servers   int `json:"servers,omitempty"`
	Addresses int `json:"addresses,omitempty"`
}

// Model is an object that represents a cloud provider location in
// the Engine Yard API.
type Model struct {
	ID         string  `json:"id,omitempty"`
	LocationID string  `json:"location_id,omitempty"`
	Limits     *Limits `json:"limits,omitempty"`

	// Timestamps
	CreatedAt  string `json:"created_at,omitempty"`
	DisabledAt string `json:"disabled_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`

	// Provider
	ProviderID int `json:"provider_id,omitempty"`

	// Parent
	ProviderLocationID string `json:"parent_id,omitempty"`
}

func (model *Model) String() string {
	if model == nil {
		return "Null location detected!"
	}
	return fmt.Sprintf("%s (Addresses: %d, Servers: %d, Provider: %d, Parent: %s)",
		model.LocationID,
		model.Limits.Addresses,
		model.Limits.Servers,
		model.ProviderID,
		model.ProviderLocationID,
	)
}
