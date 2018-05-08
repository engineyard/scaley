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

// Package users provides data types and utilities for interacting with the
// user-related endpoints on the Engine Yard Core API.
package users

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
)

// All returns the array of all users visible to the current user.
func All(api core.Client, params url.Values) []*Model {
	collection := &Collection{}

	return getUserCollection(api, collection, params)
}

// Current returns the user object associated with the provided client as well
// as an error.
//
// If errors are encountered while pulling the user information from the API,
// then the returned error is non-nil and the returned user is nil. Otherwise,
// the returned user is non-nil and the returned error is nil.
func Current(api core.Client) (*Model, error) {
	data, err := api.Get(currentUserPath(), nil)
	if err != nil {
		return nil, err
	}

	wrapper := struct {
		User *Model
	}{nil}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.User, nil
}

func currentUserPath() string {
	return "/users/current"
}

func getUserCollection(c core.Client, uc *Collection,
	params url.Values) []*Model {

	var err error

	page := 1

	all := make([]*Model, 0)
	uc.Collected = append(uc.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(uc.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(userCollectionPath(col), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				uc.Collected = col.Collected
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

func userCollectionPath(uc *Collection) string {
	return "/users"
}
