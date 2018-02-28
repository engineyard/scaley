package finders

import (
	"net/url"
	"strings"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/servers"
)

func EnvironmentForServer(api core.Client, server *servers.Model) *environments.Model {
	parts := strings.Split(server.EnvironmentURI, "/")

	params := url.Values{}
	params.Set("id", parts[len(parts)-1])

	return environments.All(api, params)[0]
}
