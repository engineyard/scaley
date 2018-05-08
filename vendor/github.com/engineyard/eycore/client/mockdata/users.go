package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/users"
)

var currentUser *users.Model
var userStore []*users.Model

func Users() []*users.Model {
	return userStore
}

func AddUser(user *users.Model) (*users.Model, error) {
	var err error

	if len(user.ID) < 0 {
		user.ID = strconv.Itoa(len(userStore) + 1)
	}

	userStore = append(userStore, user)

	return user, err
}

func SetCurrentUser(user *users.Model) {
	currentUser = user
}

func CurrentUser() *users.Model {
	return currentUser
}

func GetUsers(parts []string, params url.Values) []*users.Model {
	contenders := userStore

	for key := range params {
		if key == nameParam {
			contenders = usersWithName(contenders, params.Get(nameParam))
		}

		if key == "email" {
			contenders = usersWithEmail(contenders, params.Get("email"))
		}
	}

	return contenders
}

func GetUser(id string, parts []string, params url.Values) *users.Model {
	var ret *users.Model

	if id == "current" {
		ret = currentUser
	} else {
		contenders := GetUsers(parts, params)
		withID := usersWithID(contenders, id)

		if len(withID) > 0 {
			ret = withID[0]
		}
	}
	return ret
}

func usersWithID(contenders []*users.Model, id string) []*users.Model {
	ret := make([]*users.Model, 0)

	for _, user := range contenders {
		if user.ID == id {
			ret = append(ret, user)
		}
	}

	return ret
}

func usersWithName(contenders []*users.Model, name string) []*users.Model {
	ret := make([]*users.Model, 0)

	for _, user := range contenders {
		if user.Name == name {
			ret = append(ret, user)
		}
	}

	return ret
}

func usersWithEmail(contenders []*users.Model, email string) []*users.Model {
	ret := make([]*users.Model, 0)

	for _, user := range contenders {
		if user.Email == email {
			ret = append(ret, user)
		}
	}

	return ret
}
