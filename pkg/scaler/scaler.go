package scaler

import (
	"fmt"
	"strings"

	"github.com/engineyard/eycore/core"

	"github.com/engineyard/scaley/pkg/basher"
	"github.com/engineyard/scaley/pkg/common"
)

const (
	upFail   = "Errors occurred while starting these servers, please contact support: %s"
	downFail = "Errors occurred while stopping these servers, please contact support: %s"
)

type strategy interface {
	Upscale() error
	Downscale() error
}

type scalable interface {
	ScalingStrategy() string
	Candidates(string) []Server
	PreStop() string
}

type Server interface {
	AmazonID() string
	EngineYardID() int
	Hostname() string
}

func For(group scalable, api core.Client) strategy {
	if strings.ToLower(group.ScalingStrategy()) == "individual" {
		return newIndividual(group, api)
	}

	return newLegion(group, api)
}

func startServer(s Server, api core.Client) error {
	req, err := common.ServerReq(
		api,
		fmt.Sprintf("/servers/%d/start", s.EngineYardID()),
	)

	if err != nil {
		return err
	}

	req, err = common.WaitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}

func stopServer(s Server, api core.Client, shutdown string) error {
	if len(shutdown) > 0 {
		status := basher.Run(fmt.Sprintf("%s %s", shutdown, s.Hostname()))

		if status != 0 {
			return fmt.Errorf("The stop script for %s failed. The server has been left running so this problem can be investigated.", s.AmazonID())
		}
	}

	req, err := common.ServerReq(
		api,
		fmt.Sprintf("/servers/%d/stop", s.EngineYardID()),
	)

	if err != nil {
		return err
	}

	req, err = common.WaitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}
