package mockdata

import (
	"net/url"

	"github.com/engineyard/eycore/providers"
)

var providerStore []*providers.Model

func Providers() []*providers.Model {
	return providerStore
}

func AddProvider(provider *providers.Model) (*providers.Model, error) {
	var err error

	if provider.ID < 0 {
		provider.ID = len(providerStore) + 1
	}

	providerStore = append(providerStore, provider)

	return provider, err
}

func GetProviders(parts []string, params url.Values) []*providers.Model {
	contenders := providerStore

	if len(parts) > 1 {
		accountid := Peek(parts)
		contenders = providersForAccount(contenders, accountid)
	}

	return contenders
}

func providersForAccount(contenders []*providers.Model, accountid string) []*providers.Model {
	ret := make([]*providers.Model, 0)

	for _, provider := range contenders {
		if provider.AccountID == accountid {
			ret = append(ret, provider)
		}
	}

	return ret
}
