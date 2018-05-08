package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/accounts"
)

var account1 = &accounts.Model{
	ID:   "1",
	Name: "Account 1",
}

var account2 = &accounts.Model{
	ID:   "2",
	Name: "Account 2",
}

func setupMockAccounts() {
	accountStore = []*accounts.Model{account1, account2}
}

func TestGetAccountsWithoutParams(t *testing.T) {
	setupMockAccounts()

	result := GetAccounts(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all accounts request to return all accounts, got", result)
	}
}

func TestGetAccountsByName(t *testing.T) {
	setupMockAccounts()

	params := url.Values{}
	params.Set("name", "Account 2")

	result := GetAccounts(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	account := result[0]
	if account.ID != account2.ID {
		t.Error("Expected account 2, got", result[0])
	}

}

func TestGetAccountsByUserID(t *testing.T) {
	setupMockAccounts()

	withUser := &accounts.Model{
		ID:     "31337",
		Name:   "Owned Account",
		UserID: "12345",
	}

	accountStore = append(accountStore, withUser)

	result := GetAccounts(
		[]string{"users", "12345"},
		nil,
	)

	account := result[0]
	if account.ID != withUser.ID {
		t.Error("Expected the owned account, got", account)
	}
}

func TestGetAccount(t *testing.T) {
	setupMockAccounts()

	result := GetAccount("2", nil, nil)

	if result.ID != account2.ID {
		t.Error("Expected account 2, got", result)
	}

}
