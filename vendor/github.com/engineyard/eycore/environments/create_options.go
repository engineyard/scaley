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

import (
	"encoding/json"
)

type CreateOptions struct {
	Name         string `json:"name"`
	Framework    string `json:"framework_env,omitempty"`
	Language     string `json:"language,omitempty"`
	DeployMethod string `json:"deploy_method,omitempty"`
	Region       string `json:"region,omitempty"`
}

func NewCreateOptions() *CreateOptions {
	return &CreateOptions{}
}

func (params *CreateOptions) Body() []byte {
	var data []byte

	wrapper := struct {
		CreateOptions *CreateOptions `json:"environment,omitempty"`
	}{params}

	data, _ = json.Marshal(wrapper)

	return data
}
