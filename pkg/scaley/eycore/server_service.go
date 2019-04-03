package eycore

import (
	"fmt"
	"strings"

	"github.com/ess/eygo"

	"github.com/engineyard/scaley/pkg/scaley"
)

type ServerService struct {
	upstream *eygo.ServerService
}

func NewServerService() *ServerService {
	return &ServerService{
		eygo.NewServerService(Driver),
	}
}

func (service *ServerService) Get(provisionedID string) (*scaley.Server, error) {
	params := eygo.Params{}
	params.Set("provisioned_id", provisionedID)

	collection := service.upstream.All(params)

	if len(collection) > 1 {
		return nil, fmt.Errorf("more than one server with id %s found", provisionedID)
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("no server with id %s found", provisionedID)
	}

	s := collection[0]
	envParts := strings.Split(s.EnvironmentURL, "/")

	server := &scaley.Server{
		ID:            s.ID,
		ProvisionedID: s.ProvisionedID,
		State:         s.State,
		EnvironmentID: envParts[len(envParts)-1],
	}

	return server, nil
}

func (service *ServerService) Start(server *scaley.Server) error {

	req, err := serverReq(fmt.Sprintf("/servers/%d/start", server.ID))
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

func (service *ServerService) Stop(server *scaley.Server) error {
	return fmt.Errorf("Unimplemented")
}
