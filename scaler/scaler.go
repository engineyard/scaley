package scaler

import (
	"fmt"
	"strings"

	"github.com/engineyard/eycore/core"

	"github.com/engineyard/scaley/common"
)

type strategy interface {
	Upscale() error
	Downscale() error
}

type scalable interface {
	ScalingStrategy() string
	Candidates(string) []Server
}

type Server interface {
	AmazonID() string
	EngineYardID() int
}

func For(group scalable, api core.Client) strategy {
	if strings.ToLower(group.ScalingStrategy()) == "single" {
		return nil // newSingle(g, api)
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

func stopServer(s Server, api core.Client) error {
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
