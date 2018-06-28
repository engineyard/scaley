package eyv3

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/engineyard/eycore"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/servers"

	"github.com/engineyard/scaley/pkg/scaley"
)

type ServerService struct {
	Driver core.Client
}

func NewServerService(host string, token string) *ServerService {
	return &ServerService{Driver: eycore.NewClient(host, token)}
}

func (service *ServerService) Get(provisionedID string) (scaley.Server, error) {
	params = url.Values{}
	params.Set("provisioned_id", provisionedID)

	collection := servers.All(service.Driver, params)

	if len(collection) == 0 {
		return scaley.Server{ProvisionedID: provisionedID}, fmt.Errorf("not found")
	}

	server := collection[0]

	return scaley.Server{
		ID:            server.ID,
		ProvisionedID: server.ProvisionedID,
		State:         service.state(server),
		EnvironmentID: service.environmentID(server),
	}
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

func (service *ServerService) environmentID(server *servers.Model) int {
	id := 0

	parts := strings.Split(server.EnvironmentURI, "/")
	id = strconv.Atoi(parts[len(parts)-1])

	return id
}
