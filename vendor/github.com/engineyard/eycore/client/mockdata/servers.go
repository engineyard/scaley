package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/servers"
)

var serverStore []*servers.Model

func Servers() []*servers.Model {
	return serverStore
}

func AddServer(server *servers.Model) (*servers.Model, error) {
	var err error

	if server.ID < 0 {
		server.ID = len(serverStore) + 1
	}

	serverStore = append(serverStore, server)

	return server, err
}

func GetServers(parts []string, params url.Values) []*servers.Model {
	contenders := serverStore

	if len(parts) > 1 {
		envid := Peek(parts)
		contenders = serversForEnvironment(contenders, envid)
	}

	for key := range params {
		switch key {
		case nameParam:
			contenders = serversWithName(contenders, params.Get(nameParam))
		case "provisioned_id":
			contenders = serversWithProvisionedID(contenders, params.Get("provisioned_id"))
		case "role":
			contenders = serversWithRole(contenders, params.Get("role"))
		}
	}

	return contenders
}

func GetServer(id string, parts []string,
	params url.Values) *servers.Model {

	var ret *servers.Model

	contenders := GetServers(parts, params)
	withID := serversWithID(contenders, id)

	if len(withID) > 0 {
		ret = withID[0]
	}

	return ret
}

func serversForEnvironment(contenders []*servers.Model, envid string) []*servers.Model {
	ret := make([]*servers.Model, 0)

	for _, server := range contenders {
		expected, _ := strconv.Atoi(envid)
		if server.EnvironmentID == expected {
			ret = append(ret, server)
		}
	}

	return ret
}

func serversWithID(contenders []*servers.Model, id string) []*servers.Model {
	ret := make([]*servers.Model, 0)

	for _, server := range contenders {
		serverid, _ := strconv.Atoi(id)
		if server.ID == serverid {
			ret = append(ret, server)
		}
	}

	return ret
}

func serversWithName(contenders []*servers.Model, name string) []*servers.Model {
	ret := make([]*servers.Model, 0)

	for _, server := range contenders {
		if server.Name == name {
			ret = append(ret, server)
		}
	}

	return ret
}

func serversWithProvisionedID(contenders []*servers.Model, id string) []*servers.Model {
	ret := make([]*servers.Model, 0)

	for _, server := range contenders {
		if server.ProvisionedID == id {
			ret = append(ret, server)
		}
	}

	return ret
}

func serversWithRole(contenders []*servers.Model, role string) []*servers.Model {
	ret := make([]*servers.Model, 0)

	for _, server := range contenders {
		if server.Role == role {
			ret = append(ret, server)
		}
	}

	return ret
}
