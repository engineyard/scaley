package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/requests"
)

var requestStore []*requests.Model

func Requests() []*requests.Model {
	return requestStore
}

func GetRequests(parts []string, params url.Values) []*requests.Model {
	contenders := requestStore

	active := len(params.Get("active")) > 0
	params.Del("active")

	if active {
		contenders = requestsNotFinished(contenders)
	}

	for key := range params {
		switch key {
		case "type":
			contenders = requestsWithType(contenders, params.Get("type"))
		}
	}

	return contenders
}

func GetRequest(id string, parts []string,
	params url.Values) *requests.Model {

	var ret *requests.Model

	contenders := GetRequests(parts, params)
	withID := requestsWithID(contenders, id)

	if len(withID) > 0 {
		ret = withID[0]
	}

	return ret
}

func AddRequest(req *requests.Model) (*requests.Model, error) {
	var err error

	if len(req.ID) == 0 {
		req.ID = strconv.Itoa(len(requestStore) + 1)
	}
	requestStore = append(requestStore, req)

	return req, err
}

func requestsWithID(contenders []*requests.Model, id string) []*requests.Model {
	ret := make([]*requests.Model, 0)

	for _, request := range contenders {
		if request.ID == id {
			ret = append(ret, request)
		}
	}

	return ret
}

func requestsWithType(contenders []*requests.Model, rtype string) []*requests.Model {
	ret := make([]*requests.Model, 0)

	for _, request := range contenders {
		if request.Type == rtype {
			ret = append(ret, request)
		}
	}

	return ret
}

func requestsNotFinished(contenders []*requests.Model) []*requests.Model {
	ret := make([]*requests.Model, 0)

	for _, request := range contenders {
		if len(request.FinishedAt) == 0 {
			ret = append(ret, request)
		}
	}

	return ret
}
