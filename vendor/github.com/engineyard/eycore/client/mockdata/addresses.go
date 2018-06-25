package mockdata

import (
	"net/url"

	"github.com/engineyard/eycore/addresses"
)

var addressStore []*addresses.Model

func Addresses() []*addresses.Model {
	return addressStore
}

func AddAddress(address *addresses.Model) (*addresses.Model, error) {
	var err error

	if address.ID < 1 {
		address.ID = len(addressStore) + 1
	}

	addressStore = append(addressStore, address)

	return address, err

}

func GetAddresses(parts []string, params url.Values) []*addresses.Model {
	contenders := addressStore

	if len(parts) > 1 {
		accountid := Peek(parts)
		contenders = addressesForAccount(contenders, accountid)
	}

	for key := range params {
		switch key {
		case "location":
			contenders = addressesWithLocation(contenders, params.Get("location"))
		}
	}

	return contenders
}

func addressesForAccount(contenders []*addresses.Model, accountid string) []*addresses.Model {
	ret := make([]*addresses.Model, 0)

	for _, address := range contenders {
		if address.AccountID == accountid {
			ret = append(ret, address)
		}
	}

	return ret
}

func addressesWithLocation(contenders []*addresses.Model, location string) []*addresses.Model {
	ret := make([]*addresses.Model, 0)

	for _, address := range contenders {
		if address.Location == location {
			ret = append(ret, address)
		}
	}

	return ret
}
