package fs

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/pkg/scaley"
)

// GroupService provides the functionality around reading a group definition
// from the file system.
type GroupService struct {
	ServerService      scaley.ServerService
	EnvironmentService scaley.EnvironmentService
}

// NewGroupService takes a ServerService and an EnvironmentService, returning
// a shiny new GroupService.
func NewGroupService(servers scaley.ServerService, environments scaley.EnvironmentService) *GroupService {
	return &GroupService{
		ServerService:      servers,
		EnvironmentService: environments,
	}
}

// Get takes a group name and attempts to load that named group from the file
// system. If there are any errors in this process, they are returned along
// with the possibly unusable group.
func (service *GroupService) Get(name string) (scaley.Group, error) {
	var err error
	group := scaley.Group{Name: name}

	dir := GroupConfigs()
	file := fmt.Sprintf("%s/%s.yml", dir, name)

	if !FileExists(file) {
		return group, fmt.Errorf("No group named '%s'", name)
	}

	data, err := ReadFile(file)
	if err != nil {
		return group, err
	}

	wrapped := struct {
		Name          string   `yaml:"name"`
		Servers       []string `yaml:"scaling_servers"`
		ScalingScript string   `yaml:"scaling_script"`
		StopScript    string   `yaml:"stop_script"`
		Strategy      string   `yaml:"strategy"`
	}{Name: name}

	err = yaml.Unmarshal(data, &wrapped)
	if err != nil {
		return group, err
	}

	servers := service.servers(wrapped.Servers)
	if len(servers) == 0 {
		return group, fmt.Errorf("could not find any associated servers")
	}

	environment, err := service.environment(servers)
	if err != nil {
		return group, fmt.Errorf("could not find an associated environment")
	}

	group = scaley.Group{
		Name:          name,
		Servers:       servers,
		ScalingScript: wrapped.ScalingScript,
		StopScript:    wrapped.StopScript,
		Strategy:      wrapped.Strategy,
		Environment:   environment,
	}

	return group, nil
}

func (service *GroupService) servers(ids []string) []scaley.Server {
	result := make([]scaley.Server, 0)

	for _, id := range ids {
		if server, err := service.ServerService.Get(id); err == nil {
			result = append(result, server)
		}
	}

	return result
}

func (service *GroupService) environment(servers []scaley.Server) (scaley.Environment, error) {

	for _, server := range servers {
		if found, err := service.EnvironmentService.Get(server.EnvironmentID); err == nil {
			return found, nil
		}
	}

	return scaley.Environment{}, fmt.Errorf("environment not found")
}

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
