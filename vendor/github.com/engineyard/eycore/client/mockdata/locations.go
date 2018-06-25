package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/locations"
)

var providerLocationStore []*locations.Model

func ProviderLocations() []*locations.Model {
	return providerLocationStore
}

func AddProviderLocation(location *locations.Model) (*locations.Model, error) {
	var err error

	if len(location.ID) == 0 {
		location.ID = strconv.Itoa(len(providerLocationStore) + 1)
	}

	providerLocationStore = append(providerLocationStore, location)

	return location, err
}

func GetProviderLocations(parts []string, params url.Values) []*locations.Model {
	contenders := providerLocationStore

	if len(parts) > 1 {
		var parentid string
		parentid, parts = Pop(parts)

		switch Peek(parts) {
		case "provider-locations":
			contenders = locationsForProviderLocation(contenders, parentid)
		case "providers":
			providerid, _ := strconv.Atoi(parentid)
			contenders = locationsForProvider(contenders, providerid)
		}
	}

	return contenders
}

func locationsForProvider(contenders []*locations.Model, providerid int) []*locations.Model {
	ret := make([]*locations.Model, 0)

	for _, location := range contenders {
		if location.ProviderID == providerid {
			ret = append(ret, location)
		}
	}

	return ret
}

func locationsForProviderLocation(contenders []*locations.Model, parentid string) []*locations.Model {
	ret := make([]*locations.Model, 0)

	for _, location := range contenders {
		if location.ProviderLocationID == parentid {
			ret = append(ret, location)
		}
	}

	return ret
}
