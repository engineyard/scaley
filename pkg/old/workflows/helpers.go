// Copyright © 2018 Engine Yard, Inc.
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

package workflows

import (
	//"fmt"

	"github.com/engineyard/eycore"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/users"

	"github.com/engineyard/scaley/pkg/util"
)

type authenticated func(api core.Client, current *users.Model) error

// asCurrentUser is a helper that performs an "authenticated" function within
// the context of an authenticated Engine Yard API client and the user that is
// authenticated via said client.
func asCurrentUser(process authenticated) error {

	err := util.CheckToken()
	if err != nil {
		return err
	}

	api := eycore.NewClient(util.ApiHost(), util.Token())

	current, err := users.Current(api)
	if err != nil {
		return err
	}

	err = process(api, current)

	return err
}
