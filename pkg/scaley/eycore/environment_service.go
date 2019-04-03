package eycore

import (
	"fmt"

	"github.com/ess/eygo"

	"github.com/engineyard/scaley/pkg/scaley"
)

type EnvironmentService struct {
	upstream *eygo.EnvironmentService
}

func NewEnvironmentService() *EnvironmentService {
	return &EnvironmentService{
		eygo.NewEnvironmentService(Driver),
	}
}

func (service *EnvironmentService) Get(id string) (*scaley.Environment, error) {
	return nil, fmt.Errorf("Unimplemented")
}

func (service *EnvironmentService) Configure(env *scaley.Environment) error {
	return fmt.Errorf("Unimplemented")
}
