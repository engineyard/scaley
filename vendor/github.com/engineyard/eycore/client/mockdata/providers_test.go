package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/providers"
)

var provider1 = &providers.Model{
	ID: 1,
}

var provider2 = &providers.Model{
	ID: 2,
}

func setupMockProviders() {
	providerStore = []*providers.Model{provider1, provider2}
}

func TestGetProvidersWithoutParams(t *testing.T) {
	setupMockProviders()

	result := GetProviders(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all providers request to return all providers, got", result)
	}
}

func TestGetProvidersByAccountID(t *testing.T) {
	setupMockProviders()

	withAccount := &providers.Model{
		ID:        31337,
		AccountID: "12345",
	}

	providerStore = append(providerStore, withAccount)

	result := GetProviders(
		[]string{"accounts", "12345"},
		nil,
	)

	provider := result[0]
	if provider.ID != withAccount.ID {
		t.Error("Expected the owned provider, got", provider)
	}
}
