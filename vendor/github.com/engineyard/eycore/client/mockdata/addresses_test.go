package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/addresses"
)

var address1 = &addresses.Model{
	ID:            1,
	ProvisionedID: "1.1.1.1",
	IPAddress:     "1.1.1.1",
	AccountID:     "12345",
	Location:      "location-1",
}

var address2 = &addresses.Model{
	ID:            2,
	ProvisionedID: "2.2.2.2",
	IPAddress:     "2.2.2.2",
	AccountID:     "54321",
	Location:      "location-2",
}

func setupMockAddresss() {
	addressStore = []*addresses.Model{address1, address2}
}

func TestGetAddresssWithoutParams(t *testing.T) {
	setupMockAddresss()

	result := GetAddresses(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all addresses request to return all addresses, got", result)
	}
}

func TestGetAddresssByAccountID(t *testing.T) {
	setupMockAddresss()

	result := GetAddresses(
		[]string{"accounts", "12345"},
		nil,
	)

	address := result[0]
	if address.ID != address1.ID {
		t.Error("Expected address 1, got", address)
	}
}

func TestGetAddresssByLocation(t *testing.T) {
	setupMockAddresss()

	params := url.Values{}
	params.Set("location", "location-2")

	result := GetAddresses(
		nil,
		params,
	)

	address := result[0]
	if address.ID != address2.ID {
		t.Error("Expected address 2, got", address)
	}
}
