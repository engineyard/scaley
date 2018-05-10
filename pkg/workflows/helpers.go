package workflows

import (
	//"fmt"

	"github.com/engineyard/eycore"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/users"

	"github.com/engineyard/scaley/pkg/util"
)

type authenticated func(api core.Client, current *users.Model) error

// asCurrentUser is a helper that performs an "authenticated" function within
// the context of an authenticated Engine Yard API client and the user that is
// authenticated via said client.
func asCurrentUser(process authenticated) error {

	err := util.CheckToken()
	if err != nil {
		return err
	}

	api := eycore.NewClient(util.ApiHost(), util.Token())

	current, err := users.Current(api)
	if err != nil {
		return err
	}

	err = process(api, current)

	return err
}
