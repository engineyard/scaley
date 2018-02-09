package group

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/requests"
)

func Scale(group *Group, api core.Client, direction string) error {
	if direction == "up" {
		return upscale(group, api)
	}

	return downscale(group, api)
}

func upscale(group *Group, api core.Client) error {
	// 1. Start all scaling servers
	failures := make([]string, 0)

	for _, server := range group.ScalingServers {
		err := startServer(server, api)
		if err != nil {
			failures = append(failures, server.ID)
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf("Errors occurred while starting these servers, please contact support: %s", strings.Join(failures, ", "))
	}

	// 2. Run chef on the environment
	if err := runChef(api, group.Environment); err != nil {
		return fmt.Errorf("A Chef error occurred while upscaling the group. Please contact support.")
	}

	return nil
}

func runChef(api core.Client, environment *environments.Model) error {
	req, err := environments.Apply(api, environment, "")
	if err != nil {
		return err
	}

	req, err = waitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}

func downscale(group *Group, api core.Client) error {
	// 1. Stop all of the scaling servers
	failures := make([]string, 0)

	for _, server := range group.ScalingServers {
		err := stopServer(server, api)
		if err != nil {
			failures = append(failures, server.ID)
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf("Errors occurred while stopping these servers: %s", strings.Join(failures, ", "))
	}

	// 2. Run chef on the environment
	if err := runChef(api, group.Environment); err != nil {
		return fmt.Errorf("A Chef error occurred while upscaling the group. Please contact support.")
	}

	return nil
}

func startServer(server *Server, api core.Client) error {
	// Only act on stopped servers
	if server.Instance.State == "stopped" {
		req, err := serverReq(api, fmt.Sprintf("/servers/%d/start", server.Instance.ID))
		if err != nil {
			return err
		}

		req, err = waitFor(api, req)
		if err != nil {
			return err
		}

		if !req.Successful {
			return fmt.Errorf("%s", req.RequestStatus)
		}
	}

	return nil
}

func stopServer(server *Server, api core.Client) error {
	// Only act on servers that are not stopped
	if server.Instance.State != "stopped" {
		req, err := serverReq(api, fmt.Sprintf("/servers/%d/stop", server.Instance.ID))
		if err != nil {
			return err
		}

		req, err = waitFor(api, req)
		if err != nil {
			return err
		}

		if !req.Successful {
			return fmt.Errorf("%s", req.RequestStatus)
		}

	}

	return nil
}

func serverReq(api core.Client, path string) (*requests.Model, error) {
	params := url.Values{}

	data, err := api.Put(path, params, nil)
	if err != nil {
		return nil, fmt.Errorf("The request to PUT %s failed", path)
	}

	wrapper := struct {
		Request *requests.Model `json:"request,omitempty"`
	}{nil}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, fmt.Errorf("The API returned an invalid response when doing PUT %s", path)
	}

	return wrapper.Request, nil
}

func waitFor(api core.Client, req *requests.Model) (*requests.Model, error) {
	var err error

	ret := req

	for len(ret.FinishedAt) == 0 {
		time.Sleep(5 * time.Second)

		ret, err = requests.Refresh(api, req)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
