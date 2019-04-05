package eycore

import (
	"encoding/json"
	"time"

	"github.com/ess/eygo"
	"github.com/ess/eygo/http"
)

// Driver is an eygo.Driver instance that is used for all API operations within
// the package.
var Driver eygo.Driver

// Setup takes a url and an authentication token and sets up the package's
// Driver to interact with the EY Core API at that url with the given token.
func Setup(baseURL string, token string) {
	if Driver == nil {
		Driver, _ = http.NewDriver(baseURL, token)
	}
}

func serverReq(path string) (*eygo.Request, error) {
	response := Driver.Put(path, nil, nil)
	if response.Okay() {
		data := response.Pages[0]

		wrapper := struct {
			Request *eygo.Request `json:"request"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Request, nil
	}

	return nil, response.Error
}

func rawPost(path string) (*eygo.Request, error) {
	response := Driver.Post(path, nil, nil)
	if response.Okay() {
		data := response.Pages[0]

		wrapper := struct {
			Request *eygo.Request `json:"request"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Request, nil
	}

	return nil, response.Error

}

func waitFor(req *eygo.Request) (*eygo.Request, error) {
	var err error

	requests := eygo.NewRequestService(Driver)

	ret := req

	for len(ret.FinishedAt) == 0 {
		time.Sleep(5 * time.Second)

		ret, err = requests.Find(req.ID)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Copyright © 2019 Engine Yard, Inc.
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
