// Copyright © 2017 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/debugging"
	"github.com/engineyard/eycore/logger"
)

// NewCoreAPI instantiates and returns a CoreAPI with the provided API host and
// token.
func NewCoreAPI(host string, token string) *CoreAPI {
	if host == "" {
		host = "api.engineyard.com"
	}

	BaseURL := url.URL{
		Scheme: "https",
		Host:   host,
	}

	return &CoreAPI{
		&http.Client{
			Timeout: 20 * time.Second,
		},
		BaseURL,
		token,
	}
}

// CoreAPI is an object that provides the low-level mechanism for interacting
// with the Engine Yard Core API.
type CoreAPI struct {
	*http.Client
	BaseURL url.URL
	Token   string
}

// Get performs a GET operation for the given path and URL query against the
// Engine Yard Core API. It returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the raw response body and the
// error is nil.
func (api *CoreAPI) Get(path string, params url.Values) ([]byte, error) {
	body, err := api.makeRequest("GET", api.constructRequestURL(path, params), nil)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post performs a POST operation for the given path and URL query against the
// Engine Yard Core API, sending the provided data along in its payload. It
// returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the raw response body and the
// error is nil.
func (api *CoreAPI) Post(path string, params url.Values, data core.Body) ([]byte, error) {

	body, err := api.makeRequest("POST", api.constructRequestURL(path, params), data)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Put performs a PUT operation for the given path and URL query against the
// Engine Yard Core API, sending the provided data along in its payload. It
// returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the raw response body and the
// error is nil.
func (api *CoreAPI) Put(path string, params url.Values, data core.Body) ([]byte, error) {
	body, err := api.makeRequest("PUT", api.constructRequestURL(path, params), data)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Delete performs a DELETE operation for the given path and URL query against
// the Engine Yard Core API. It returns a byte array as well as an error.
//
// If any errors are present in the process, the returned error is non-nil and
// the body is nil. Otherwise, the body is the raw response body and the
// error is nil.
func (api *CoreAPI) Delete(path string, params url.Values) ([]byte, error) {
	body, err := api.makeRequest("DELETE", api.constructRequestURL(path, params), nil)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (api *CoreAPI) setRequestHeaders(request *http.Request) {
	request.Header.Add("X-Ey-Token", api.Token)
	request.Header.Set("Accept", "application/vnd.engineyard.v3+json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Host", "api.engineyard.com")
	request.Header.Set("User-Agent", "eycore/0.1.0 (https://www.engineyard.com)")
}

func extractErrors(response *http.Response, body []byte) error {
	err := errors.New(response.Status)
	var r map[string]interface{}
	jsonErr := json.Unmarshal(body, &r)

	if jsonErr != nil {
		return err
	}

	if coreErrors, ok := r["errors"]; ok {
		// If we get a useful error message from Core, we create a less generic Error with a user-friendly printout
		// of the Core error.
		coreError := core.NewError("Could not process request. The following issues were identified:\n")

		for _, e := range coreErrors.([]interface{}) {
			coreError.ErrorString += e.(string) + "\n"
		}

		return coreError
	}

	return err
}

func (api *CoreAPI) makeRequest(method string, requestURL string, data core.Body) ([]byte, error) {

	body := make([]byte, 0)

	if data != nil {
		body = data.Body()
	}

	if debugging.Enabled() {
		logger.Info("CoreAPI.makeRequest",
			fmt.Sprintf("Handling %s for %s Body: (%s)",
				method,
				requestURL,
				string(body)))
	}

	request, err := http.NewRequest(method, requestURL, bytes.NewReader(body))
	if err != nil {
		if debugging.Enabled() {
			logger.Error("CoreAPI.makeRequest",
				fmt.Sprintf("Creating the request failed - %s", err.Error()))
		}
		return nil, err
	}

	api.setRequestHeaders(request)

	response, err := api.Do(request)

	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(response.Body)
	//response.Body.Close()
	/*fmt.Println("response body:", string(body))*/

	if err != nil {
		return nil, err
	}

	if debugging.Enabled() {
		logger.Info("CoreAPI.makeRequest",
			fmt.Sprintf("API Response Body: %s", string(body)))
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, extractErrors(response, body)
	}

	if debugging.Enabled() {
		fmt.Println()
	}

	return body, nil
}

func (api *CoreAPI) constructRequestURL(path string, params url.Values) string {
	requestURL := url.URL{
		Scheme:   api.BaseURL.Scheme,
		Host:     api.BaseURL.Host,
		Path:     path,
		RawQuery: processParams(params),
	}

	return requestURL.String()
}
