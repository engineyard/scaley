package fs

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/pkg/scaley"
)

// GroupService is a service that knows how to interact with groups on the
// file system.
type GroupService struct{}

// NewGroupService returns a new instance of GroupService.
func NewGroupService() *GroupService {
	return &GroupService{}
}

// Get takes a group name, attempts to read that group's definition from the
// file system, and returns both a group and an error. If there are issues
// along the way, the error is populated and the group is nil. Otherwise, the
// group is populated from the values in the associated group file and the error
// is nil.
func (s *GroupService) Get(name string) (*scaley.Group, error) {
	var err error

	g := &scaley.Group{Name: name, Strategy: "legion"}

	dir := GroupConfigs()
	file := fmt.Sprintf("%s/%s.yml", dir, name)

	if !FileExists(file) {
		return nil, fmt.Errorf("No group named '%s'", name)
	}

	data, err := ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
