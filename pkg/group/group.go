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

package group

import (
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/pkg/common"
	"github.com/engineyard/scaley/pkg/finders"
)

type Group struct {
	Name             string    `yaml:"name"`
	PermanentServers []*Server `yaml:"permanent_servers"`
	ScalingServers   []*Server `yaml:"scaling_servers"`
	ScalingScript    string    `yaml:"scaling_script"`
	StopScript       string    `yaml:"stop_script"`
	Strategy         string    `yaml:"strategy"`
	Environment      *environments.Model
}

func ByName(api core.Client, name string) (*Group, error) {
	var err error
	group := &Group{
		Name:     name,
		Strategy: "legion",
	}

	dir := common.GroupConfigs()
	file := fmt.Sprintf("%s/%s.yml", dir, name)

	if !common.FileExists(file) {
		return nil, fmt.Errorf("No group named '%s'", name)
	}

	data, err := common.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, group)
	if err != nil {
		return nil, err
	}

	for _, server := range group.ScalingServers {
		server.Instance = finders.FindServer(api, server.ID)
	}

	for _, server := range group.ScalingServers {
		if server.Instance != nil {
			group.Environment = finders.EnvironmentForServer(
				api,
				group.ScalingServers[0].Instance,
			)
			break
		}
	}

	return group, nil
}
