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

// Package flavors provides data types and utilities for interacting with the
// server-flavor-related endpoints on the Engine Yard Core API.
package flavors

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
)

// AllForAccount returns an array that contains all flavors that are
// visible to the current user and associated with the given account.
func AllForAccount(api core.Client, account *accounts.Model, params url.Values) []*Model {

	collection := &Collection{AccountID: account.ID}

	return getFlavorCollection(api, collection, params)
}

func flavorCollectionPath(collection *Collection) string {
	return "/flavors"
}

func getFlavorCollection(c core.Client, collection *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	collection.Collected = append(collection.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	if len(params.Get("account")) == 0 {
		params.Set("account", collection.AccountID)
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(collection.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(flavorCollectionPath(collection), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				collection.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	return all
}
