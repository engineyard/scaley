package mockdata

import (
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/flavors"
	"github.com/engineyard/eycore/locations"
)

var flavorStore []*flavors.Model

func Flavors() []*flavors.Model {
	return flavorStore
}

func AddFlavor(flavor *flavors.Model) (*flavors.Model, error) {
	var err error

	if len(flavor.ID) == 0 {
		flavor.ID = strconv.Itoa(len(flavorStore) + 1)
	}

	flavorStore = append(flavorStore, flavor)

	return flavor, err
}

func GetFlavors(parts []string, params url.Values) []*flavors.Model {
	var account *accounts.Model
	var region *locations.Model

	for _, a := range accountStore {
		if a.ID == params.Get("account") {
			account = a
			break
		}
	}

	providers := providersForAccount(providerStore, account.ID)
	locations := locationsForProvider(providerLocationStore, providers[0].ID)

	for _, l := range locations {
		if l.LocationID == params.Get("location") {
			region = l
			break
		}
	}

	return flavorsForProviderLocation(flavorStore, region.ID)
}

func flavorsForProviderLocation(contenders []*flavors.Model, locationid string) []*flavors.Model {
	ret := make([]*flavors.Model, 0)

	for _, flavor := range contenders {
		if flavor.ProviderLocationID == locationid {
			ret = append(ret, flavor)
		}
	}

	return ret
}
