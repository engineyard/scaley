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

package flavors

import (
	"fmt"
)

// Model is an object that represents a server flavor in the Engine Yard
// API.
type Model struct {
	ID              string `json:"id,omitempty"`
	APIName         string `json:"api_name,omitempty"`
	Description     string `json:"description,omitempty"`
	Dedicated       bool   `json:"dedicated,omitempty"`
	VolumeOptimized bool   `json:"volume_optimized,omitempty"`
	Architecture    int    `json:"architecture,omitempty"`
	Name            string `json:"name,omitempty"`

	// These are not real associations, but are added as a convenience
	AccountID          string `json:"account_id,omitempty"`
	ProviderLocationID string `json:"provider_location_id,omitempty"`
}

func (model *Model) String() string {
	return fmt.Sprintf("%s (%s)", model.APIName, model.Description)
}
