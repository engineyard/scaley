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

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/addresses"
	"github.com/engineyard/eycore/client/mockdata"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/debugging"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/flavors"
	"github.com/engineyard/eycore/locations"
	"github.com/engineyard/eycore/logger"
	"github.com/engineyard/eycore/providers"
	"github.com/engineyard/eycore/requests"
	"github.com/engineyard/eycore/servers"
	"github.com/engineyard/eycore/users"
)

const (
	envsPart    = "environments"
	serversPart = "servers"
)

// MockAPI can be used in place of a real API client in cases in which you
// want to test the behavior of an application that uses eycore. It does not
// behave exactly like the actual upstream Engine Yard API, but is basically
// meant to be "good enough for government work."
type MockAPI struct {
	BaseURL url.URL
	Token   string
}

// NewMockAPI returns a MockAPI instance configured with the provided host and
// token.
func NewMockAPI(host string, token string) *MockAPI {
	if host == "" {
		host = "api.engineyard.com"
	}

	BaseURL := url.URL{
		Scheme: "https",
		Host:   host,
	}

	ret := &MockAPI{
		BaseURL: BaseURL,
		Token:   token,
	}

	return ret
}

// Get simulates a GET response for the provided path and URL params against
// the Engine Yard Core API using a pre-populated collection of mock data. It
// returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the simulated raw response body and
// the error is nil.
//
// Additionally, in an effort to try to emulate the real API, params may be
// partially or entirely ignored, depending on the request in question.
func (api *MockAPI) Get(path string, params url.Values) ([]byte, error) {
	id, model, parts := api.idModelParts(path)

	if len(params.Get("page")) > 0 {
		page, err := strconv.Atoi(params.Get("page"))
		if err != nil {
			page = 0
		}

		if page != 1 {
			ret := []byte(`{"` + model + `":[]}`)
			return ret, nil
		}
	}

	return api.getModel(model, id, parts, params)
}

// Post simulates a POST operation for the given path and URL query against the
// Engine Yard Core API using a pre-populated collection of mock data, sending
// the provided data along in its payload. It returns a byte array as well as
// an error.
//
// If any errors are present in the process, the returned error is non-nil, the
// body is nil, and the data set does not change. Otherwise, the body is the
// raw response body, the error is nil, and the provided data is typically added
// to the mock data set.
//
// Additionally, in an effort to try to emulate the real API, params may be
// partially or entirely ignored, depending on the request in question.
func (api *MockAPI) Post(path string, params url.Values, data core.Body) ([]byte, error) {
	id, model, parts := api.idModelParts(path)

	//fmt.Println("id:", id, "model:", model, "parts:", parts)
	return api.postModel(id, model, parts, params, data)
}

// Put simulates a PUT operation for the given path and URL query against the
// Engine Yard Core API using a pre-populated collection of mock data, sending
// the provided data along in its payload. It returns a byte array as well as
// an error.
//
// If any errors are present in the process, the returned error is non-nil, the
// body is nil, and the data set does not change. Otherwise, the body is the
// raw response body, the error is nil, and the mock data set is typically
// altered to reflect the provided data.
//
// Additionally, in an effort to try to emulate the real API, params may be
// partially or entirely ignored, depending on the request in question.
func (api *MockAPI) Put(path string, params url.Values, data core.Body) ([]byte, error) {
	return nil, errors.New("PUT not implemented")
}

// Delete simulates a DELETE response for the provided path and URL params
// against the Engine Yard Core API using a pre-populated collection of mock
// data. It returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the simulated raw response body and
// the error is nil.
//
// Additionally, in an effort to try to emulate the real API, params may be
// partially or entirely ignored, depending on the request in question.
func (api *MockAPI) Delete(path string, params url.Values) ([]byte, error) {
	id, model, parts := api.idModelParts(path)

	return api.deleteModel(id, model, parts, params)
}

func (api *MockAPI) isModelPart(part string) bool {
	models := []string{
		"servers",
		"environments",
		"accounts",
		"users",
	}

	for _, model := range models {
		if part == model {
			return true
		}
	}

	return false
}

type usersWrapper struct {
	Users []*users.Model `json:"users,omitempty"`
}

type userWrapper struct {
	User *users.Model `json:"user,omitempty"`
}

type accountsWrapper struct {
	Accounts []*accounts.Model `json:"accounts,omitempty"`
}

type addressesWrapper struct {
	Addresses []*addresses.Model `json:"addresses,omitempty"`
}

type providersWrapper struct {
	Providers []*providers.Model `json:"providers,omitempty"`
}

type providerLocationsWrapper struct {
	ProviderLocations []*locations.Model `json:"provider_locations,omitempty"`
}

type flavorsWrapper struct {
	Flavors []*flavors.Model `json:"flavors,omitempty"`
}

type accountWrapper struct {
	Account *accounts.Model `json:"account,omitempty"`
}

type environmentsWrapper struct {
	Environments []*environments.Model `json:"environments,omitempty"`
}

type environmentWrapper struct {
	Environment *environments.Model `json:"environment,omitempty"`
}

type serversWrapper struct {
	Servers []*servers.Model `json:"servers,omitempty"`
}

type serverWrapper struct {
	Server *servers.Model `json:"server,omitempty"`
}

type requestWrapper struct {
	Request *requests.Model `json:"request,omitempty"`
}

type requestsWrapper struct {
	Requests []*requests.Model `json:"requests,omitempty"`
}

func (api *MockAPI) deleteModel(id string, model string, parts []string, params url.Values) ([]byte, error) {
	var ret []byte
	var err error

	if len(id) > 0 {
		switch model {
		case "environments":
			req := &requests.Model{Type: "destroy_environment"}
			req, err = mockdata.AddRequest(req)

			if err == nil {
				ret, err = json.Marshal(&requestWrapper{Request: req})
			}
		}
	}
	return ret, err
}

func (api *MockAPI) postModel(id string, model string, parts []string, params url.Values, body core.Body) ([]byte, error) {
	var data []byte
	var ret []byte
	var err error

	if body == nil {
		data = make([]byte, 0)
	} else {
		data = body.Body()
	}

	if len(id) > 0 {
		switch id {
		case "apply":
			req := &requests.Model{Type: "configure_environment"}
			req, err = mockdata.AddRequest(req)

			if err == nil {
				ret, err = json.Marshal(&requestWrapper{Request: req})
			}
		case "boot":
			req := &requests.Model{Type: "start_environment"}
			req, err = mockdata.AddRequest(req)

			if err == nil {
				ret, err = json.Marshal(&requestWrapper{Request: req})
			}
		case "tokens":
			if len(os.Getenv("MOCK_EYCORE_UNAUTHORIZED")) > 0 {
				err = errors.New(`{"errors":{"You are not authorized"}}`)
			} else {
				ret = []byte(`{"api_token":"tokatokatoken"}`)
			}
		}
	} else {
		switch model {
		case envsPart:
			env := &environments.Model{}
			err = json.Unmarshal(data, &environmentWrapper{Environment: env})

			if err == nil {
				env, err = mockdata.AddEnvironment(env)
				if err == nil {
					ret, err = json.Marshal(&environmentWrapper{Environment: env})
				}
			}
		case serversPart:
			req := &requests.Model{Type: "provision_server"}
			req, err = mockdata.AddRequest(req)

			if err == nil {
				ret, err = json.Marshal(&requestWrapper{Request: req})
			}
		default:
			err = errors.New(IllegalOperation)
		}
	}

	return ret, err
}

func (api *MockAPI) getModel(model string, id string, parts []string,
	params url.Values) ([]byte, error) {

	if debugging.Enabled() {
		logger.Info("MockAPI.getModel",
			fmt.Sprintf("model: %s, id: %s, parts [%s]", model, id, strings.Join(parts, ", ")))
	}

	var err error
	ret := make([]byte, 0)

	if len(id) > 0 {
		// We got something that looks like an ID, so try to determine the
		// model

		switch model {
		case "users":
			user := mockdata.GetUser(id, parts, params)
			if user == nil {
				err = core.NewError("User not found")
			} else {
				ret, err = json.Marshal(&userWrapper{User: user})
			}
		case "accounts":
			account := mockdata.GetAccount(id, parts, params)
			if account == nil {
				err = core.NewError("Account not found")
			} else {
				ret, err = json.Marshal(&accountWrapper{Account: account})
			}
		case envsPart:
			environment := mockdata.GetEnvironment(id, parts, params)
			if environment == nil {
				err = core.NewError("Environment not found")
			} else {
				ret, err = json.Marshal(&environmentWrapper{Environment: environment})
			}
		case serversPart:
			server := mockdata.GetServer(id, parts, params)
			if server == nil {
				err = core.NewError("Server not found")
			} else {
				ret, err = json.Marshal(&serverWrapper{Server: server})
			}
		case "requests":
			request := mockdata.GetRequest(id, parts, params)
			if request == nil {
				err = core.NewError("Request not found")
			} else {
				ret, err = json.Marshal(&requestWrapper{Request: request})

				// Set the request as finished so subsequent get terminates a wait
				if len(request.FinishedAt) == 0 {
					request.FinishedAt = "just now"
					request.Successful = true
				}
			}

		default:
			// Okay, so we didn't really get an ID

			switch id {
			case "addresses":
				parts = append(parts, model)
				addresses := mockdata.GetAddresses(parts, params)
				ret, err = json.Marshal(&addressesWrapper{Addresses: addresses})
			case "providers":
				parts = append(parts, model)
				providers := mockdata.GetProviders(parts, params)
				ret, err = json.Marshal(&providersWrapper{Providers: providers})
			case "locations":
				parts = append(parts, model)
				locations := mockdata.GetProviderLocations(parts, params)
				ret, err = json.Marshal(&providerLocationsWrapper{ProviderLocations: locations})
			case "provider-locations":
				parts = append(parts, model)
				locations := mockdata.GetProviderLocations(parts, params)
				ret, err = json.Marshal(&providerLocationsWrapper{ProviderLocations: locations})
			case "flavors":
				flavors := mockdata.GetFlavors(parts, params)
				ret, err = json.Marshal(&flavorsWrapper{Flavors: flavors})

			case "requests":
				parts = append(parts, model)
				requests := mockdata.GetRequests(parts, params)
				ret, err = json.Marshal(&requestsWrapper{Requests: requests})

			}
		}

	} else {
		switch model {
		case "users":
			users := mockdata.GetUsers(parts, params)
			ret, err = json.Marshal(&usersWrapper{Users: users})
		case "accounts":
			accounts := mockdata.GetAccounts(parts, params)
			ret, err = json.Marshal(&accountsWrapper{Accounts: accounts})
		case "environments":
			environments := mockdata.GetEnvironments(parts, params)
			ret, err = json.Marshal(&environmentsWrapper{Environments: environments})
		case "servers":
			servers := mockdata.GetServers(parts, params)
			ret, err = json.Marshal(&serversWrapper{Servers: servers})
		}
	}

	return ret, err
}

func (api *MockAPI) pathParts(path string) []string {
	return strings.Split(path, "/")
}

func (api *MockAPI) modelParts(path string) (string, []string) {
	return mockdata.Pop(api.pathParts(path))
}

func (api *MockAPI) idModelParts(path string) (string, string, []string) {
	current, parts := api.modelParts(path)

	var model string
	var id string

	if api.isModelPart(current) {
		model = current
	} else {
		id = current
		model, parts = mockdata.Pop(parts)
	}

	return id, model, parts
}
