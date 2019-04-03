package eycore

import (
	"github.com/ess/eygo"
	"github.com/ess/eygo/http"
)

var Driver eygo.Driver

func Setup(baseURL string, token string) {
	if Driver == nil {
		Driver, _ = http.NewDriver(baseURL, token)
	}
}
