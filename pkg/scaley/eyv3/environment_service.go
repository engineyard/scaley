package eyv3

import (
	"encoding/json"
	"fmt"

	"github.com/engineyard/eycore/environments"

	"github.com/engineyard/scaley/pkg/scaley"
)

// EnvironmentService provides the functionality for retrieving environment data
// from the Engine Yard Core-v3 API.
type EnvironmentService struct {
	Driver Reader
}

// NewEnvironmentService instantiates a new EnvironmentService with the given
// API reader
func NewEnvironmentService(driver Reader) *EnvironmentService {
	return &EnvironmentService{Driver: driver}
}

// Get retrieves the specified environment from the upstream API. If there are
// errors along the way, an error is returned. Otherwise, a scaley Environment
// is returned.
func (service *EnvironmentService) Get(ID string) (scaley.Environment, error) {
	data, err := service.Driver.Get("environments/"+ID, nil)
	if err != nil {
		return scaley.Environment{ID: ID}, fmt.Errorf("not found")
	}

	wrapper := struct {
		Environment *environments.Model `json:"environment,omitempty"`
	}{}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return scaley.Environment{ID: ID}, fmt.Errorf("received invalid upstream environment data")
	}

	model := wrapper.Environment

	environment := scaley.Environment{
		ID:   ID,
		Name: model.Name,
	}

	return environment, nil
}
