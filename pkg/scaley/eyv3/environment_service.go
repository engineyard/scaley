package eyv3

import (
	"encoding/json"
	"fmt"

	"github.com/engineyard/eycore/environments"

	"github.com/engineyard/scaley/pkg/scaley"
)

// EnvironmentService provides the functionality for retrieving environment data
// from the Engine Yard Core-v3 API.
type EnvironmentService struct {
	Driver Reader
}

// NewEnvironmentService instantiates a new EnvironmentService with the given
// API reader
func NewEnvironmentService(driver Reader) *EnvironmentService {
	return &EnvironmentService{Driver: driver}
}

// Get retrieves the specified environment from the upstream API. If there are
// errors along the way, an error is returned. Otherwise, a scaley Environment
// is returned.
func (service *EnvironmentService) Get(ID string) (scaley.Environment, error) {
	data, err := service.Driver.Get("environments/"+ID, nil)
	if err != nil {
		return scaley.Environment{ID: ID}, fmt.Errorf("not found")
	}

	wrapper := struct {
		Environment *environments.Model `json:"environment,omitempty"`
	}{}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return scaley.Environment{ID: ID}, fmt.Errorf("received invalid upstream environment data")
	}

	model := wrapper.Environment

	environment := scaley.Environment{
		ID:   ID,
		Name: model.Name,
	}

	return environment, nil
}

// Copyright © 2018 Engine Yard, Inc.
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
