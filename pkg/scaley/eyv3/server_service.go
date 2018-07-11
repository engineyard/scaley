package eyv3

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/engineyard/eycore/servers"

	"github.com/engineyard/scaley/pkg/scaley"
)

// ServerService provides the functionality for retrieving server data
// from the Engine Yard Core-v3 API.
type ServerService struct {
	Driver Client
}

// NewServerService instantiates a new ServerService with the given reader.
func NewServerService(driver Client) *ServerService {
	return &ServerService{Driver: driver}
}

// Get retrieves the specified server from the upstream API. If there are any
// errors along the way, an error is returned. Otherwise, a proper scaley
// server is returned.
func (service *ServerService) Get(provisionedID string) (scaley.Server, error) {
	params := url.Values{}
	params.Set("provisioned_id", provisionedID)

	collection := servers.All(service.Driver, params)

	if len(collection) == 0 {
		return scaley.Server{ProvisionedID: provisionedID}, fmt.Errorf("not found")
	}

	model := collection[0]

	server := scaley.Server{
		ID:            model.ID,
		ProvisionedID: model.ProvisionedID,
		State:         service.state(model),
		EnvironmentID: service.environmentID(model),
	}

	return server, nil
}

func (service *ServerService) state(server *servers.Model) int {
	state := 0

	switch server.State {
	case "stopped":
		state = scaley.Stopped
	case "running":
		state = scaley.Running
	default:
		state = scaley.Unknown
	}

	return state
}

func (service *ServerService) environmentID(server *servers.Model) string {
	parts := strings.Split(server.EnvironmentURI, "/")

	return parts[len(parts)-1]
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
