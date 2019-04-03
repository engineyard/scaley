package fs

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/pkg/scaley"
)

type GroupService struct{}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (s *GroupService) Get(name string) (*scaley.Group, error) {
	var err error

	g := struct {
		Name           string   `yaml:"name"`
		ScalingServers []string `yaml:"scaling_servers"`
		ScalingScript  string   `yaml:"scaling_script"`
		StopScript     string   `yaml:"stop_script"`
		Strategy       string   `yaml:"strategy"`
	}{Name: name, Strategy: "legion"}

	dir := GroupConfigs()
	file := fmt.Sprintf("%s/%s.yml", dir, name)

	if !FileExists(file) {
		return nil, fmt.Errorf("No group named '%s'", name)
	}

	data, err := ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &g)
	if err != nil {
		return nil, err
	}

	group := &scaley.Group{
		Name:           g.Name,
		ScalingServers: g.ScalingServers,
		ScalingScript:  g.ScalingScript,
		StopScript:     g.StopScript,
		Strategy:       g.Strategy,
	}

	return group, nil
}
