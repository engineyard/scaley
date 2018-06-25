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

// Package client provides a low-level interface for interacting with the
// Engine Yard Core API.
package client

import (
	"net/url"
	"sort"
	"strings"
)

const (
	// NotFound indicates that the Resource upon which the client is operating
	// was not found.
	NotFound = "404"

	// Forbidden indicates that the authenticated user is not allowed to access
	// a given Resource or API endpoint.
	Forbidden = "403"

	// IllegalOperation is used to express that an attempted operation on the
	// upstream API is not allowed.
	IllegalOperation = "This operation is not allowed in the client."
)

func processParams(params url.Values) string {
	var queryParts []string

	if params != nil {
		keys := make([]string, 0)

		for k := range params {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			//for k, v := range params {
			queryParts = append(queryParts,
				url.QueryEscape(key)+"="+url.QueryEscape(params.Get(key)))
		}
	}

	return strings.Join(queryParts, "&")
}
