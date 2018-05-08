package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/environments"
)

var environmentStore []*environments.Model

func Environments() []*environments.Model {
	return environmentStore
}

func AddEnvironment(env *environments.Model) (*environments.Model, error) {
	var err error

	if env.ID <= 0 {
		env.ID = len(environmentStore) + 1
	}

	environmentStore = append(environmentStore, env)

	return env, err
}

func GetEnvironments(parts []string, params url.Values) []*environments.Model {
	contenders := environmentStore

	if len(parts) > 1 {
		accountid := Peek(parts)
		contenders = environmentsForAccount(contenders, accountid)
	}

	for key := range params {
		switch key {
		case nameParam:
			contenders = environmentsWithName(contenders, params.Get(nameParam))
		}
	}

	return contenders
}

func GetEnvironment(id string, parts []string,
	params url.Values) *environments.Model {

	var ret *environments.Model

	contenders := GetEnvironments(parts, params)
	withID := environmentsWithID(contenders, id)

	if len(withID) > 0 {
		ret = withID[0]
	}

	return ret
}

func environmentsForAccount(contenders []*environments.Model, accountid string) []*environments.Model {
	ret := make([]*environments.Model, 0)

	for _, environment := range contenders {
		if environment.AccountID == accountid {
			ret = append(ret, environment)
		}
	}

	return ret
}

func environmentsWithID(contenders []*environments.Model, id string) []*environments.Model {
	ret := make([]*environments.Model, 0)

	for _, environment := range contenders {
		envid, _ := strconv.Atoi(id)
		if environment.ID == envid {
			ret = append(ret, environment)
		}
	}

	return ret
}

func environmentsWithName(contenders []*environments.Model, name string) []*environments.Model {
	ret := make([]*environments.Model, 0)

	for _, environment := range contenders {
		if environment.Name == name {
			ret = append(ret, environment)
		}
	}

	return ret
}
