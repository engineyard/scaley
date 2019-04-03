package eycore

import (
	"encoding/json"
	"time"

	"github.com/ess/eygo"
	"github.com/ess/eygo/http"
)

var Driver eygo.Driver

func Setup(baseURL string, token string) {
	if Driver == nil {
		Driver, _ = http.NewDriver(baseURL, token)
	}
}

func serverReq(path string) (*eygo.Request, error) {
	response := Driver.Put(path, nil, nil)
	if response.Okay() {
		data := response.Pages[0]
		req := &eygo.Request{}
		err := json.Unmarshal(data, req)
		if err != nil {
			return nil, err
		}

		return req, nil
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
