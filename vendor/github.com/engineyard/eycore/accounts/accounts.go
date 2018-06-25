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

// Package accounts provides data types and utilities for interacting with the
// account-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/accounts
package accounts

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
	"github.com/engineyard/eycore/users"
)

// All returns an array that contains all accounts visible to the current
// user.
//
// This account collection can be refined via the parameters listed at
// http://developer.engineyard.com/accounts#list-all-accounts
func All(api core.Client, params url.Values) []*Model {
	collection := &Collection{}

	return getAccountCollection(api, collection, params)
}

// AllForUser returns an array that contains all accounts that are visible to
// the current user and associated with the given user.
//
// This account collection can be refined via the parameters listed at
// http://developer.engineyard.com/accounts#list-all-accounts-for-a-given-user
func AllForUser(api core.Client, user *users.Model,
	params url.Values) []*Model {

	collection := &Collection{UserID: user.ID}

	return getAccountCollection(api, collection, params)
}

func accountCollectionPath(ac *Collection) string {
	if ac.UserID != "" {
		return fmt.Sprintf("/users/%s/accounts", ac.UserID)
	}

	return "/accounts"
}

func getAccountCollection(c core.Client, ac *Collection,
	params url.Values) []*Model {

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
		if data, err := c.Get(accountCollectionPath(ac), params); err == nil {
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

	for _, account := range all {
		account.UserID = ac.UserID
	}

	return all
}
