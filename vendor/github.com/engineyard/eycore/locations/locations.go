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

// Package locations provides data types and utilities for interacting with the
// provider-location-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/provider_locations
package locations

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
	"github.com/engineyard/eycore/providers"
)

// AllForProvider returns an array that contains all locations that are
// associated with the given provider.
func AllForProvider(api core.Client, provider *providers.Model,
	params url.Values) []*Model {

	collection := &Collection{ProviderID: provider.ID}

	return getProviderLocationCollection(api, collection, params)
}

// Children returns an array that contains all locations that are
// associated with the given location.
func Children(api core.Client, location *Model,
	params url.Values) []*Model {

	collection := &Collection{ProviderLocationID: location.ID}

	return getProviderLocationCollection(api, collection, params)
}

func getProviderLocationCollection(c core.Client, pc *Collection, params url.Values) []*Model {
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
		if data, err := c.Get(providerLocationCollectionPath(pc), params); err == nil {
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

	for _, providerLocation := range all {
		providerLocation.ProviderID = pc.ProviderID
	}

	return all
}

func providerLocationCollectionPath(collection *Collection) string {
	if len(collection.ProviderLocationID) > 0 {
		return fmt.Sprintf("/provider-locations/%s/provider-locations", collection.ProviderLocationID)
	}

	return fmt.Sprintf("/providers/%d/locations", collection.ProviderID)
}
