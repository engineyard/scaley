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

package applications

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/paging"
)

func All(api core.Client, params url.Values) []*Model {
	collection := &Collection{}

	return getApplicationCollection(api, collection, params)
}

func AllForEnvironment(api core.Client, environment *environments.Model,
	params url.Values) []*Model {

	collection := &Collection{EnvironmentID: environment.ID}

	return getApplicationCollection(api, collection, params)
}

func getApplicationCollection(c core.Client, ac *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	ac.Collected = append(ac.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(ac.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(applicationCollectionPath(ac), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				ac.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	for _, application := range all {
		application.EnvironmentID = ac.EnvironmentID
	}

	return all
}

func applicationCollectionPath(ac *Collection) string {
	if ac.EnvironmentID > 0 {
		return fmt.Sprintf("/environments/%d/applications", ac.EnvironmentID)
	}

	return "/applications"
}
