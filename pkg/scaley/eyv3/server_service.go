package eyv3

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/servers"

	"github.com/engineyard/scaley/pkg/scaley"
)

// ServerService provides the functionality for retrieving server data
// from the Engine Yard Core-v3 API.
type ServerService struct {
	Driver Reader
}

// NewServerService instantiates a new ServerService with the given reader.
func NewServerService(driver Reader) *ServerService {
	return &ServerService{Driver: driver}
}

// Get retrieves the specified server from the upstream API. If there are any
// errors along the way, an error is returned. Otherwise, a proper scaley
// server is returned.
func (service *ServerService) Get(provisionedID string) (scaley.Server, error) {
	params := url.Values{}
	params.Set("provisioned_id", provisionedID)

	collection := servers.All(service.Driver, params)

	if len(collection) == 0 {
		return scaley.Server{ProvisionedID: provisionedID}, fmt.Errorf("not found")
	}

	model := collection[0]

	server := scaley.Server{
		ID:            model.ID,
		ProvisionedID: model.ProvisionedID,
		State:         service.state(model),
		EnvironmentID: service.environmentID(model),
	}

	return server, nil
}

func (service *ServerService) state(server *servers.Model) int {
	state := 0

	switch server.State {
	case "stopped":
		state = 1
	case "running":
		state = 2
	}

	return state
}

func (service *ServerService) environmentID(server *servers.Model) string {
	parts := strings.Split(server.EnvironmentURI, "/")

	return parts[len(parts)-1]
}
