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

package mockdata

import (
	"encoding/json"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/addresses"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/flavors"
	"github.com/engineyard/eycore/locations"
	"github.com/engineyard/eycore/providers"
	"github.com/engineyard/eycore/requests"
	"github.com/engineyard/eycore/servers"
	"github.com/engineyard/eycore/users"
)

const (
	nameParam = "name"
)

func Pop(a []string) (string, []string) {
	return a[len(a)-1], a[:len(a)-1]
}

func Peek(a []string) string {
	return a[len(a)-1]
}

type dataSet struct {
	CurrentUser       *users.Model          `json:"current_user,omitempty"`
	Users             []*users.Model        `json:"users,omitempty"`
	Accounts          []*accounts.Model     `json:"accounts,omitempty"`
	Environments      []*environments.Model `json:"environments,omitempty"`
	Servers           []*servers.Model      `json:"servers,omitempty"`
	Requests          []*requests.Model     `json:"requests,omitempty"`
	Addresses         []*addresses.Model    `json:"addresses,omitempty"`
	Providers         []*providers.Model    `json:"providers,omitempty"`
	ProviderLocations []*locations.Model    `json:"provider_locations,omitempty"`
	Flavors           []*flavors.Model      `json:"flavors,omitempty"`
}

func Seed(raw []byte) {
	ret := &dataSet{}

	json.Unmarshal(raw, ret)

	currentUser = ret.CurrentUser
	userStore = ret.Users
	accountStore = ret.Accounts
	environmentStore = ret.Environments
	serverStore = ret.Servers
	requestStore = ret.Requests
	addressStore = ret.Addresses
	providerStore = ret.Providers
	providerLocationStore = ret.ProviderLocations
	flavorStore = ret.Flavors
}
