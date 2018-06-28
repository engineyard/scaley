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

package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/requests"
)

func ServerReq(api core.Client, path string) (*requests.Model, error) {
	params := url.Values{}

	data, err := api.Put(path, params, nil)
	if err != nil {
		return nil, fmt.Errorf("The request to PUT %s failed", path)
	}

	wrapper := struct {
		Request *requests.Model `json:"request,omitempty"`
	}{nil}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, fmt.Errorf("The API returned an invalid response when doing PUT %s", path)
	}

	return wrapper.Request, nil
}

func WaitFor(api core.Client, req *requests.Model) (*requests.Model, error) {
	var err error

	ret := req

	for len(ret.FinishedAt) == 0 {
		time.Sleep(5 * time.Second)

		ret, err = requests.Refresh(api, req)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
