package finders

import (
	"net/url"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/servers"
)

func FindServer(api core.Client, serverID string) *servers.Model {
	params := url.Values{}
	params.Set("provisioned_id", serverID)

	collection := servers.All(api, params)

	if len(collection) == 0 {
		return nil
	}

	return collection[0]
}
