package eycore

import (
	"encoding/json"
	"time"

	"github.com/ess/eygo"
	"github.com/ess/eygo/http"
)

// Driver is an eygo.Driver instance that is used for all API operations within
// the package.
var Driver eygo.Driver

// Setup takes a url and an authentication token and sets up the package's
// Driver to interact with the EY Core API at that url with the given token.
func Setup(baseURL string, token string) {
	if Driver == nil {
		Driver, _ = http.NewDriver(baseURL, token)
	}
}

func serverReq(path string) (*eygo.Request, error) {
	response := Driver.Put(path, nil, nil)
	if response.Okay() {
		data := response.Pages[0]

		wrapper := struct {
			Request *eygo.Request `json:"request"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Request, nil
	}

	return nil, response.Error
}

func rawPost(path string) (*eygo.Request, error) {
	response := Driver.Post(path, nil, nil)
	if response.Okay() {
		data := response.Pages[0]

		wrapper := struct {
			Request *eygo.Request `json:"request"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return nil, err
		}

		return wrapper.Request, nil
	}

	return nil, response.Error

}

func waitFor(req *eygo.Request) (*eygo.Request, error) {
	var err error

	requests := eygo.NewRequestService(Driver)

	ret := req

	for len(ret.FinishedAt) == 0 {
		time.Sleep(5 * time.Second)

		ret, err = requests.Find(req.ID)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
