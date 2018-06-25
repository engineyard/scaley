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

package environments

import "fmt"

// Model is an object that represents an environment in the Engine Yard
// API.
type Model struct {
	// Standard Environment Fields
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Environment Details
	DatabaseStack      string `json:"database_stack,omitempty"`
	DeployMethod       string `json:"deploy_method,omitempty"`
	FrameworkEnv       string `json:"framework_env,omitempty"`
	InternalPrivateKey string `json:"internal_private_key,omitempty"`
	InternalPublicKey  string `json:"internal_public_key,omitempty"`
	Language           string `json:"language,omitempty"`
	Region             string `json:"region,omitempty"`
	ReleaseLabel       string `json:"release_label,omitempty"`
	ServiceLevel       string `json:"service_level,omitempty"`
	ServicePlan        string `json:"service_plan,omitempty"`
	StackName          string `json:"stack_name,omitempty"`
	UserName           string `json:"username,omitempty"`

	// Timestamps
	CreatedAt string `json:"created_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	// The account associated with the environment. This is tagged for marshaling
	// simply as a convenience and is not actually sent in this form from the
	// API.
	AccountID string `json:"account_id,omitempty"`
}

// String returns a plain text represenation of the Environment.
func (model *Model) String() string {
	return fmt.Sprintf("%s (%d)", model.Name, model.ID)
}
