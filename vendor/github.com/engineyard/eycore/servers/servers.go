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

// Package servers provides data types and utilities for interacting with the
// server-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/servers
package servers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/paging"
)

// All returns an array that contains all servers visible to the current user.
//
// This server collection can be refined via the parameters listed at
// http://developer.engineyard.com/servers#list-servers
func All(api core.Client, params url.Values) []*Model {
	collection := &Collection{}

	return getServerCollection(api, collection, params)
}

// AllForEnvironment returns an array that contains all servers that are
// associated with the given environment.
//
// This server collection can be refined via the parameters listed at
// http://developer.engineyard.com/servers#list-servers
func AllForEnvironment(api core.Client, environment *environments.Model,
	params url.Values) []*Model {

	collection := &Collection{EnvironmentID: environment.ID}

	return getServerCollection(api, collection, params)
}

func getServerCollection(c core.Client, sc *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	sc.Collected = append(sc.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(sc.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(serverCollectionPath(sc), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				sc.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	for _, server := range all {
		server.EnvironmentID = sc.EnvironmentID
	}

	return all
}

func serverCollectionPath(sc *Collection) string {
	if sc.EnvironmentID > 0 {
		return fmt.Sprintf("/environments/%d/servers", sc.EnvironmentID)
	}

	return "/servers"
}
