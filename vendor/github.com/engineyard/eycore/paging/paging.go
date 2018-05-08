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

// Package paging provides helpers for pulling paginated results from the
// Engine Yard Core API.
package paging

import (
	"strconv"
	"time"
)

const (
	// PerPage specifies the number of results that we want per page when
	// retrieving records from the API.
	PerPage = "100"

	// PollTime specifies the amount of time to wait between API calls when
	// pulling multiple pages.
	PollTime = 5 * time.Second

	// FinalPage is used to indicate that there are no further results forthcoming
	// from the API.
	FinalPage = "final page"
)

// MaxResults is the maximum number of results that can be returned from a
// single paginated call to the API.
func MaxResults() int {
	// TODO: Figure out what to do if Atoi errors here. Technically, that should
	// NEVER happen because we're getting the data from a const string, but it
	// is reported by gas, so it's worth investigating.
	max, _ := strconv.Atoi(PerPage)
	return max
}

// Page is used to determine if the API has further results for a given request.
// If the passed error is a FinalPage, a nil error is returned. Otherwise,
// the error that was passed in is returned.
func Page(err error) error {
	if err != nil && err.Error() == FinalPage {
		return nil
	}

	return err
}
