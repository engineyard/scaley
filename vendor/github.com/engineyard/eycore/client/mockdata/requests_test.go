package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/requests"
)

var request1 = &requests.Model{
	ID:   "1",
	Type: "Sausages",
}

var request2 = &requests.Model{
	ID:         "2",
	Type:       "Pigeons",
	FinishedAt: "some point",
}

func setupMockRequests() {
	requestStore = []*requests.Model{request1, request2}
}

func TestGetRequestsWithoutParams(t *testing.T) {
	setupMockRequests()

	result := GetRequests(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all requests request to return all requests, got", result)
	}
}

func TestGetRequestsByType(t *testing.T) {
	setupMockRequests()

	params := url.Values{}
	params.Set("type", "Pigeons")

	result := GetRequests(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	request := result[0]
	if request.ID != request2.ID {
		t.Error("Expected request 2, got", result[0])
	}

}

func TestGetRequestsActiveOnly(t *testing.T) {
	setupMockRequests()

	params := url.Values{}
	params.Set("active", "true")

	result := GetRequests(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	request := result[0]
	if request.ID != request1.ID {
		t.Error("Expected the unfinished request, got", request)
	}
}

func TestGetRequest(t *testing.T) {
	setupMockRequests()

	result := GetRequest("2", nil, nil)

	if result.ID != request2.ID {
		t.Error("Expected request 2, got", result)
	}

}
