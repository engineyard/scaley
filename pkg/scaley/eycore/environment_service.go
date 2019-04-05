package eycore

import (
	"fmt"
	"strconv"

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

	params := eygo.Params{}
	params.Set("id", id)

	collection := service.upstream.All(params)

	if len(collection) > 1 {
		return nil, fmt.Errorf("more than one environment with id %s found", id)
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("no environment with id %s found", id)
	}

	e := collection[0]

	environment := &scaley.Environment{
		ID:   strconv.Itoa(e.ID),
		Name: e.Name,
	}

	return environment, nil
}

func (service *EnvironmentService) Configure(env *scaley.Environment) error {
	req, err := rawPost(fmt.Sprintf("environments/%s/apply", env.ID))
	if err != nil {
		return err
	}

	req, err = waitFor(req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}
