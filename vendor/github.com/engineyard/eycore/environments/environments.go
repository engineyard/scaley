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

// Package environments provides data types and utilities for interacting with
// the environment-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/environments
package environments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
	"github.com/engineyard/eycore/requests"
)

// All returns an array that contains all environments visible to the current
// user.
func All(api core.Client, params url.Values) []*Model {
	collection := &Collection{}

	return getEnvironmentCollection(api, collection, params)
}

// AllForAccount returns an array that contains all environments that are
// visible to the current user and associated with the given account.
func AllForAccount(api core.Client, account *accounts.Model,
	params url.Values) []*Model {

	collection := &Collection{AccountID: account.ID}

	return getEnvironmentCollection(api, collection, params)
}

// Apply sets up a background task on the API that runs any existing Chef
// cookbooks for the environment. A "run type" is required for this operation,
// and this run type maps to either the default, main, or custom cookbook
// behaviort. Apply returns a request instance that one can use to track
// the progress of the Chef run.
func Apply(api core.Client, environment *Model, runType string) (*requests.Model, error) {
	var ret *requests.Model
	var err error

	body := &ApplyOptions{}

	if len(runType) > 0 {
		body.SetType(runType)
	}

	data, err := api.Post(applyEnvironmentPath(environment), nil, body)
	if err == nil {
		wrapper := struct {
			Request *requests.Model `json:"request,omitempty"`
		}{nil}

		err = json.Unmarshal(data, &wrapper)

		if err == nil {
			ret = wrapper.Request
		}
	}

	return ret, err
}

// Boot creates a background task on the API that provisions and configures
// the servers defined for the environment. This requires the construction
// of a params.BootEnvironment, and it returns a request instance that one
// can use to track the boot progress.
func Boot(api core.Client, environment *Model,
	bootParams *BootOptions) (*requests.Model, error) {

	var ret *requests.Model
	var err error

	//body := bootParams.Body()
	data, err := api.Post(bootEnvironmentPath(environment), nil, bootParams)
	if err == nil {
		wrapper := struct {
			Request *requests.Model `json:"request,omitempty"`
		}{nil}

		err = json.Unmarshal(data, &wrapper)

		if err == nil {
			ret = wrapper.Request
		}
	}

	return ret, err
}

// Create is used to persist a new environment model to the upstream API.
// If there are any errors in this process, a nil environment instance and an
// error are returned. Otherwise, the persisted environment instance is
// returned with a nil error.
func Create(api core.Client, account *accounts.Model,
	environment *Model) (*Model, error) {

	var ret *Model
	var err error

	if len(environment.Name) == 0 {
		return nil, errors.New("Environment creation requires a name")
	}

	p := &CreateOptions{}

	p.Name = environment.Name
	p.Framework = environment.FrameworkEnv
	p.Language = environment.Language
	p.DeployMethod = environment.DeployMethod
	p.Region = environment.Region

	ec := &Collection{AccountID: account.ID}

	data, err := api.Post(environmentCollectionPath(ec), nil, p)
	if err == nil {
		wrapper := struct {
			Environment *Model `json:"environment,omitempty"`
		}{nil}

		err = json.Unmarshal(data, &wrapper)
		if err == nil {
			ret = wrapper.Environment
		}
	}

	return ret, err
}

// Destroy creates a background task on the API to deprovision and delete
// the given environment. If there are any errors in the process, a nil
// request instance is returned along with the error. Otherwise, a request
// instance that one can use to track destroy progress is returned along
// with a nil error.
func Destroy(api core.Client, environment *Model) (*requests.Model, error) {

	var ret *requests.Model
	var err error

	data, err := api.Delete(deleteEnvironmentPath(environment), nil)
	if err == nil {
		wrapper := struct {
			Request *requests.Model `json:"request,omitempty"`
		}{nil}

		err = json.Unmarshal(data, &wrapper)

		if err == nil {
			ret = wrapper.Request
		}
	}

	return ret, err
}

func applyEnvironmentPath(e *Model) string {
	return fmt.Sprintf("/environments/%d/apply", e.ID)
}

func bootEnvironmentPath(e *Model) string {
	return fmt.Sprintf("/environments/%d/boot", e.ID)
}

func deleteEnvironmentPath(e *Model) string {
	return fmt.Sprintf("/environments/%d", e.ID)
}

func environmentCollectionPath(ec *Collection) string {
	if ec.AccountID != "" {
		return fmt.Sprintf("/accounts/%s/environments", ec.AccountID)
	}

	return "/environments"
}

func getEnvironmentCollection(c core.Client, ec *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	ec.Collected = append(ec.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(ec.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(environmentCollectionPath(ec), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				ec.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	for _, env := range all {
		env.AccountID = ec.AccountID
	}

	return all
}
