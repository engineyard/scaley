package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
)

var accountStore []*accounts.Model

func Accounts() []*accounts.Model {
	return accountStore
}

func AddAccount(account *accounts.Model) (*accounts.Model, error) {
	var err error

	if len(account.ID) == 0 {
		account.ID = strconv.Itoa(len(accountStore) + 1)
	}

	accountStore = append(accountStore, account)

	return account, err

}

func GetAccounts(parts []string, params url.Values) []*accounts.Model {
	contenders := accountStore

	if len(parts) > 1 {
		userid := Peek(parts)
		contenders = accountsForUser(contenders, userid)
	}

	for key := range params {
		switch key {
		case nameParam:
			contenders = accountsWithName(contenders, params.Get(nameParam))
		}
	}

	return contenders
}

func GetAccount(id string, parts []string, params url.Values) *accounts.Model {
	var ret *accounts.Model

	contenders := GetAccounts(parts, params)
	withID := accountsWithID(contenders, id)

	if len(withID) > 0 {
		ret = withID[0]
	}

	return ret
}

func accountsForUser(accts []*accounts.Model, userid string) []*accounts.Model {
	ret := make([]*accounts.Model, 0)

	for _, account := range accts {
		if account.UserID == userid {
			ret = append(ret, account)
		}
	}

	return ret
}

func accountsWithID(accts []*accounts.Model, id string) []*accounts.Model {
	ret := make([]*accounts.Model, 0)

	for _, account := range accts {
		if account.ID == id {
			ret = append(ret, account)
		}
	}

	return ret
}

func accountsWithName(accts []*accounts.Model, name string) []*accounts.Model {
	ret := make([]*accounts.Model, 0)

	for _, account := range accts {
		if account.Name == name {
			ret = append(ret, account)
		}
	}

	return ret
}
