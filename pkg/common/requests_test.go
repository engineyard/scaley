package common

import (
	"testing"
	"time"

	"github.com/ess/fakey"

	"github.com/engineyard/eycore/requests"
)

func includesPath(paths []string, query string) bool {
	for _, path := range paths {
		if path == query {
			return true
		}
	}

	return false
}

func TestServerReq(t *testing.T) {
	t.Run("it performs a PUT request for the given path", func(t *testing.T) {
		api := &fakey.Client{}
		path := "/some/incredibly/fake/path"

		original := api.Requests("put")

		ServerReq(api, path)

		current := api.Requests("put")

		if len(original) == len(current) {
			t.Errorf("No PUT requests were made")
		}

		if !includesPath(current, path) {
			t.Errorf("No PUT call was made to %s", path)
		}
	})

	t.Run("when the API returns an error", func(t *testing.T) {
		api := &fakey.Client{}

		_, err := ServerReq(api, "path/1")
		if err != nil && err.Error() != "The request to PUT path/1 failed" {
			t.Errorf("Expected a general API error to be returned")
		}

	})

	t.Run("when the JSON can't be parsed", func(t *testing.T) {
		api := &fakey.Client{}
		api.AddResponse("put", "{request:}")

		_, err := ServerReq(api, "path/2")
		if err != nil && err.Error() != "The API returned an invalid response when doing PUT path/2" {
			t.Errorf("Expected an invalid response error, got '%s'", err.Error())
		}
	})

	t.Run("when the API call succeeds", func(t *testing.T) {
		api := &fakey.Client{}
		api.AddResponse("put", `{"request":{}}`)

		req, err := ServerReq(api, "path/3")
		if err != nil {
			t.Errorf("Expected a successful call not to return an error")
		}

		if req == nil {
			t.Errorf("Expected a successful call to return an eycore Request")
		}
	})
}

func TestWaitFor(t *testing.T) {
	t.Run("when the request is valid", func(t *testing.T) {
		t.Run("and is already finished", func(t *testing.T) {
			api := &fakey.Client{}
			request := &requests.Model{FinishedAt: "now"}

			start := time.Now()
			result, err := WaitFor(api, request)
			finish := time.Now()

			if finish.Sub(start).Seconds() >= 5 {
				t.Errorf("Expected a completed request to immediately return")
			}

			if result == nil {
				t.Errorf("Expected to receive a valid request object")
			}

			if err != nil {
				t.Errorf("Expected no errors")
			}
		})

		t.Run("but is not finished", func(t *testing.T) {
			api := &fakey.Client{}
			api.AddResponse("get", `{"request":{"finished_at":"now"}}`)

			request := &requests.Model{}

			start := time.Now()
			result, err := WaitFor(api, request)
			finish := time.Now()

			elapsed := finish.Sub(start).Seconds()

			if elapsed < 5 || elapsed >= 6 {
				t.Errorf("Expected to only wait 5 seconds for refresh")
			}

			if result == nil {
				t.Errorf("Expected to receive a valid request object")
			}

			if err != nil {
				t.Errorf("Expected no errors")
			}
		})
	})

	t.Run("when the request fails to refresh", func(t *testing.T) {
		api := &fakey.Client{}
		request := &requests.Model{}

		result, err := WaitFor(api, request)

		if err == nil {
			t.Errorf("Expected to receive an error")
		}

		if result != nil {
			t.Errorf("Expected to receive no request object")
		}

	})
}
