package eycore

import (
	"fmt"
	"strings"

	"github.com/ess/eygo"

	"github.com/engineyard/scaley/pkg/scaley"
)

// ServerService is a service that knows how to interact with servers via the
// EY Core API.
type ServerService struct {
	upstream *eygo.ServerService
}

// NewServerService returns a new instance of ServerService.
func NewServerService() *ServerService {
	return &ServerService{
		eygo.NewServerService(Driver),
	}
}

// Get takes a server's IaaS ID as a string and queries the upstream API for
// the server details, returning both the server and an error. If there are
// issues along the way, the error is populated and the server is nil.
// Otherwise, the server is populated and the error is nil.
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
		Hostname:      s.PrivateHostname,
		EnvironmentID: envParts[len(envParts)-1],
	}

	return server, nil
}

// Start takes a server and attempts to start it via the upstream API. If there
// are issues along the way, an error is returned. Otherwise, nil is returned.
func (service *ServerService) Start(server *scaley.Server) error {

	req, err := serverReq(fmt.Sprintf("servers/%d/start", server.ID))
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

// Stop takes a server and attempts to stop it via the upstream API. If there
// are issues along the way, an error is returned. Otherwise, nil is returned.
func (service *ServerService) Stop(server *scaley.Server) error {
	req, err := serverReq(fmt.Sprintf("servers/%d/stop", server.ID))
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
