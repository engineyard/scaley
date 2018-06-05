package fakey

import (
	"fmt"
	"net/url"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/logger"
)

type Client struct {
	requests  map[string][]string
	responses *ResponseCollection
}

func (api *Client) setup() {
	if api.responses == nil {
		api.responses = &ResponseCollection{}
	}

	if api.requests == nil {
		api.requests = make(map[string][]string)
	}
}

func (api *Client) Requests(method string) []string {
	var requests []string

	api.setup()

	requests = append(requests, api.requests[method]...)

	return requests
}

func (api *Client) AddResponse(method string, path string, response string) {
	api.setup()

	api.responses.Add(method, path, response)
}

func (api *Client) handle(method string, path string) ([]byte, error) {
	api.setup()

	api.requests[method] = append(api.requests[method], path)

	logger.Info("fakey", fmt.Sprintf("method: %s, path: %s", method, path))

	response, err := api.responses.Consume(method, path)
	if err != nil {
		logger.Error("fakey", err.Error())
		return nil, err
	}

	logger.Info("fakey", fmt.Sprintf("response: %s", response))

	return []byte(response), nil
}

func (api *Client) Get(path string, params url.Values) ([]byte, error) {
	return api.handle("get", path+api.processParams(params))
}

func (api *Client) Post(path string, params url.Values, body core.Body) ([]byte, error) {
	return api.handle("post", path+api.processParams(params))
}

func (api *Client) Put(path string, params url.Values, body core.Body) ([]byte, error) {
	return api.handle("put", path+api.processParams(params))
}

func (api *Client) Delete(path string, params url.Values) ([]byte, error) {
	return api.handle("delete", path+api.processParams(params))
}

func (api *Client) processParams(params url.Values) string {
	if len(params) > 0 {
		return "?" + params.Encode()
	}

	return ""
}
