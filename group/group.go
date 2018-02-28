package group

import (
	"fmt"
	"io/ioutil"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/common"
	"github.com/engineyard/scaley/finders"
)

type Group struct {
	Name             string    `yaml:"name"`
	PermanentServers []*Server `yaml:"permanent_servers"`
	ScalingServers   []*Server `yaml:"scaling_servers"`
	ScalingScript    string    `yaml:"scaling_script"`
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

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, group)
	if err != nil {
		return nil, err
	}

	allServers := append(group.PermanentServers, group.ScalingServers...)

	for _, server := range allServers {
		server.Instance = finders.FindServer(api, server.ID)
	}

	group.Environment = finders.EnvironmentForServer(
		api,
		group.PermanentServers[0].Instance,
	)

	return group, nil
}
