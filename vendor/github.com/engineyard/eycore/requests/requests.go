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

// Package requests provides data types and utilities for interacting with the
// request-related endpoints on the Engine Yard Core API.
//
// See the following API documentation for more information:
// http://developer.engineyard.com/requests
package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/paging"
	"github.com/ess/mockable"
)

// AllByEnvironmentID returns an array that contains all requests associated
// with the environment specified by the given environment ID.
//
// The resulting collection of requests can be filtered by the params specified
// at http://developer.engineyard.com/requests#list-requests
func AllByEnvironmentID(api core.Client, environmentID int,
	params url.Values) []*Model {

	collection := &Collection{EnvironmentID: environmentID}

	return getRequestCollection(api, collection, params)
}

// Refresh gets the most recent version of the provided request from the API,
// returning a request and an error.
//
// If an error is encountered while refreshing the request, the returned request
// is nil and the returned error is non-nil. Otherwise, the returned error is
// nil and the returned request is the most recent version of the given request
// from the API.
func Refresh(api core.Client, request *Model) (*Model, error) {
	data, err := api.Get(requestPath(request), nil)
	if err != nil {
		return nil, err
	}

	wrapper := struct {
		Request *Model `json:"request,omitempty"`
	}{nil}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Request, nil
}

// Wait polls the API for the status of the given request until either the
// request completes or the maximum wait time has been exceeded. When it is
// finished waiting, it returns a request, the duration of time waited, and
// an error.
//
// If errors were encountered while waiting, the returned request is nil and the
// returned error is non-nil. Otherwise, the returned error is nil and the
// returned request is the most recent version of the request from the API.
func Wait(api core.Client, request *Model, maxTime time.Duration) (*Model, time.Duration, error) {
	var err error
	var tries int
	var elapsed time.Duration

	pollTime := paging.PollTime

	ret := request
	tries = calculateRequestAttempts(maxTime)

	for len(ret.FinishedAt) == 0 && tries > 0 {
		ret, err = Refresh(api, request)

		// Halt the process and return the error if it is present
		if err != nil {
			request = nil
			break
		}

		if len(request.FinishedAt) == 0 {
			tries--

			// Avoid sleeping if we know we're out of tries
			if tries < 1 {
				break
			}

			if !mockable.Mocked() {
				// Sleep so as to not overload the upstream API
				time.Sleep(pollTime)
			}

			// Update the elapsed time
			elapsed = elapsed + pollTime
		}
	}

	return ret, elapsed, err
}

func calculateRequestAttempts(max time.Duration) int {
	tries := int(max / paging.PollTime)

	// We should always return at least 1 attempt, and we ignore partial attempts
	if tries < 1 {
		tries++
	}

	return tries
}

func getRequestCollection(c core.Client, rc *Collection, params url.Values) []*Model {
	var err error

	page := 1
	all := make([]*Model, 0)
	rc.Collected = append(rc.Collected, &Model{})

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", paging.PerPage)

	for paging.Page(err) == nil && len(rc.Collected) > 0 {
		params.Set("page", strconv.Itoa(page))
		col := &Collection{}
		if data, err := c.Get(requestCollectionPath(rc), params); err == nil {
			if jsonErr := json.Unmarshal(data, col); jsonErr == nil {
				rc.Collected = col.Collected
				all = append(all, col.Collected...)
				if len(col.Collected) < paging.MaxResults() {
					break
				}
			}
		}
		page = page + 1
	}

	for _, request := range all {
		request.EnvironmentID = rc.EnvironmentID
	}

	return all
}

func requestCollectionPath(rc *Collection) string {
	if rc.EnvironmentID > 0 {
		return fmt.Sprintf("/environments/%d/requests", rc.EnvironmentID)
	}

	return "/requests"
}

func requestPath(request *Model) string {
	return fmt.Sprintf("/requests/%s", request.ID)
}
