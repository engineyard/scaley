package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/servers"
)

var server1 = &servers.Model{
	ID:            1,
	Name:          "Server 1",
	ProvisionedID: "i-00000001",
	Role:          "app_master",
}

var server2 = &servers.Model{
	ID:            2,
	Name:          "Server 2",
	ProvisionedID: "i-00000002",
	Role:          "db_master",
}

func setupMockServers() {
	serverStore = []*servers.Model{server1, server2}
}

func TestGetServersWithoutParams(t *testing.T) {
	setupMockServers()

	result := GetServers(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all servers request to return all servers, got", result)
	}
}

func TestGetServersByName(t *testing.T) {
	setupMockServers()

	params := url.Values{}
	params.Set("name", "Server 2")

	result := GetServers(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	server := result[0]
	if server.ID != server2.ID {
		t.Error("Expected server 2, got", result[0])
	}

}

func TestGetServersByProvisionedID(t *testing.T) {
	setupMockServers()

	params := url.Values{}
	params.Set("provisioned_id", "i-00000002")

	result := GetServers(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	server := result[0]
	if server.ID != server2.ID {
		t.Error("Expected server 2, got", result[0])
	}

}

func TestGetServersByRole(t *testing.T) {
	setupMockServers()

	params := url.Values{}
	params.Set("role", "db_master")

	result := GetServers(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	server := result[0]
	if server.ID != server2.ID {
		t.Error("Expected server 2, got", result[0])
	}

}

func TestGetServersByEnvironmentID(t *testing.T) {
	setupMockServers()

	withEnvironment := &servers.Model{
		ID:            31337,
		Name:          "Owned Server",
		ProvisionedID: "i-00000003",
		EnvironmentID: 12345,
	}

	serverStore = append(serverStore, withEnvironment)

	result := GetServers(
		[]string{"environments", "12345"},
		nil,
	)

	server := result[0]
	if server.ID != withEnvironment.ID {
		t.Error("Expected the owned server, got", server)
	}
}

func TestGetServer(t *testing.T) {
	setupMockServers()

	result := GetServer("2", nil, nil)

	if result.ID != server2.ID {
		t.Error("Expected server 2, got", result)
	}

}
