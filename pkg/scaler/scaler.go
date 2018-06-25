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

package scaler

import (
	"fmt"
	"strings"

	"github.com/engineyard/eycore/core"

	"github.com/engineyard/scaley/pkg/basher"
	"github.com/engineyard/scaley/pkg/common"
)

const (
	upFail   = "Errors occurred while starting these servers, please contact support: %s"
	downFail = "Errors occurred while stopping these servers, please contact support: %s"
)

type strategy interface {
	Upscale() error
	Downscale() error
}

type scalable interface {
	ScalingStrategy() string
	Candidates(string) []Server
	PreStop() string
}

type Server interface {
	AmazonID() string
	EngineYardID() int
	Hostname() string
}

func For(group scalable, api core.Client) strategy {
	if strings.ToLower(group.ScalingStrategy()) == "individual" {
		return newIndividual(group, api)
	}

	return newLegion(group, api)
}

func startServer(s Server, api core.Client) error {
	req, err := common.ServerReq(
		api,
		fmt.Sprintf("/servers/%d/start", s.EngineYardID()),
	)

	if err != nil {
		return err
	}

	req, err = common.WaitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}

func stopServer(s Server, api core.Client, shutdown string) error {
	if len(shutdown) > 0 {
		status := basher.Run(fmt.Sprintf("%s %s", shutdown, s.Hostname()))

		if status != 0 {
			return fmt.Errorf("The stop script for %s failed. The server has been left running so this problem can be investigated.", s.AmazonID())
		}
	}

	req, err := common.ServerReq(
		api,
		fmt.Sprintf("/servers/%d/stop", s.EngineYardID()),
	)

	if err != nil {
		return err
	}

	req, err = common.WaitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}
