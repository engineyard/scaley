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

// Package eycore provides the models and low-level clients that allow for
// interaction with the Engine Yard Core API.
//
// See the following API documention for more information:
// http://developer.engineyard.com
package eycore

import (
	"github.com/engineyard/eycore/client"
	"github.com/engineyard/eycore/core"
	"github.com/ess/mockable"
)

// NewClient returns a new Core API client that is initialized with the provided
// API host and user token.
//
// If mocking is enabled, the resulting client is a mocked client that doesn't
// perform any actual HTTP calls, but instead relies on the existence of a mock
// data set.
//
// Otherwise, it is a client that interacts with the real Engine Yard Core API.
func NewClient(host string, token string) core.Client {
	if mockable.Mocked() {
		return client.NewMockAPI(host, token)
	}

	return client.NewCoreAPI(host, token)
}
