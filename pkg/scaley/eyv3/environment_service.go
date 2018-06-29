package eyv3

import (
	"encoding/json"
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"

	"github.com/engineyard/scaley/pkg/scaley"
)

type EnvironmentService struct {
	Driver core.Client
}

func NewEnvironmentService(driver core.Client) *EnvironmentService {
	return &EnvironmentService{Driver: driver}
}

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
