package eycore

import (
	"fmt"
	"strconv"

	"github.com/ess/debuggable"
	"github.com/ess/eygo"

	"github.com/engineyard/scaley/v2/pkg/scaley"
)

// EnvironmentService is a service that knows how to interact with Engine Yard
// Cloud environments via the EY Core API.
type EnvironmentService struct {
	upstream *eygo.EnvironmentService
}

// NewEnvironmentService returns a new instance of EnvironmentService.
func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{
		eygo.NewEnvironmentService(Driver),
	}
}

// Get takes an environment ID as a string and returns the associated
// environment and an error. If there are issues along the way, the error is
// populated and the environment is nil. Otherwise, the environment is populated
// and the error is nil.
func (service *EnvironmentService) Get(id string) (*scaley.Environment, error) {

	params := eygo.Params{}
	params.Set("id", id)

	collection := service.upstream.All(params)

	if len(collection) > 1 {
		return nil, fmt.Errorf("more than one environment with id %s found", id)
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("no environment with id %s found", id)
	}

	e := collection[0]

	environment := &scaley.Environment{
		ID:   strconv.Itoa(e.ID),
		Name: e.Name,
	}

	return environment, nil
}

// Configure takes an environment and attempts to reconfigure it. If there are
// issues on the upstream API or Chef service, an error is returned. Otherwise,
// nil is returned.
func (service *EnvironmentService) Configure(env *scaley.Environment) error {
	req, err := rawPost(fmt.Sprintf("environments/%s/apply", env.ID))
	if err != nil {
		return err
	}

	req, err = waitFor(req)
	if err != nil {
		if debuggable.Enabled() {
			fmt.Println("[scaley debug] waitFor returned an error:", err)
		}
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}

// Copyright © 2019 Engine Yard, Inc.
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
