package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/locations"
)

var providerLocation1 = &locations.Model{
	ID:                 "1",
	LocationID:         "ProviderLocation 1",
	ProviderID:         12345,
	ProviderLocationID: "54321",
}

var providerLocation2 = &locations.Model{
	ID:                 "2",
	LocationID:         "ProviderLocation 2",
	ProviderID:         67890,
	ProviderLocationID: "09876",
}

func setupMockProviderLocations() {
	providerLocationStore = []*locations.Model{providerLocation1, providerLocation2}
}

func TestGetProviderLocationsWithoutParams(t *testing.T) {
	setupMockProviderLocations()

	result := GetProviderLocations(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all providerLocations request to return all providerLocations, got", result)
	}
}

func TestGetProviderLocationsByProviderID(t *testing.T) {
	setupMockProviderLocations()

	result := GetProviderLocations(
		[]string{"providers", "12345"},
		nil,
	)

	providerLocation := result[0]
	if providerLocation.ID != providerLocation1.ID {
		t.Error("Expected location 1, got", providerLocation)
	}
}

func TestGetProviderLocationsByProviderLocationID(t *testing.T) {
	setupMockProviderLocations()

	result := GetProviderLocations(
		[]string{"provider-locations", "09876"},
		nil,
	)

	providerLocation := result[0]
	if providerLocation.ID != providerLocation2.ID {
		t.Error("Expected location 2, got", providerLocation)
	}
}
