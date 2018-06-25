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

// Package providers provides data types and utilities for interacting with the
// provider-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/providers
package providers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
)

// AllForAccount returns an array that contains all providers associated with
// the given account.
func AllForAccount(api core.Client, a *accounts.Model,
	params url.Values) [](*Model) {

	collection := &Collection{AccountID: a.ID}

	return getProviderCollection(api, collection, params)
}

func getProviderCollection(c core.Client, pc *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	pc.Collected = append(pc.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(pc.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(providerCollectionPath(pc), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				pc.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	for _, provider := range all {
		provider.AccountID = pc.AccountID
	}

	return all
}

func providerCollectionPath(pc *Collection) string {
	if pc.AccountID != "" {
		return fmt.Sprintf("/accounts/%s/providers", pc.AccountID)
	}

	return "/providers"
}
