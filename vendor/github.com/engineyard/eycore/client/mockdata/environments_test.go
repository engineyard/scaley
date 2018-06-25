package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/environments"
)

var environment1 = &environments.Model{
	ID:   1,
	Name: "Environment 1",
}

var environment2 = &environments.Model{
	ID:   2,
	Name: "Environment 2",
}

func setupMockEnvironments() {
	environmentStore = []*environments.Model{environment1, environment2}
}

func TestGetEnvironmentsWithoutParams(t *testing.T) {
	setupMockEnvironments()

	result := GetEnvironments(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all environments request to return all environments, got", result)
	}
}

func TestGetEnvironmentsByName(t *testing.T) {
	setupMockEnvironments()

	params := url.Values{}
	params.Set("name", "Environment 2")

	result := GetEnvironments(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	environment := result[0]
	if environment.ID != environment2.ID {
		t.Error("Expected environment 2, got", result[0])
	}

}

func TestGetEnvironmentsByAccountID(t *testing.T) {
	setupMockEnvironments()

	withAccount := &environments.Model{
		ID:        31337,
		Name:      "Owned Environment",
		AccountID: "12345",
	}

	environmentStore = append(environmentStore, withAccount)

	result := GetEnvironments(
		[]string{"accounts", "12345"},
		nil,
	)

	environment := result[0]
	if environment.ID != withAccount.ID {
		t.Error("Expected the owned environment, got", environment)
	}
}

func TestGetEnvironment(t *testing.T) {
	setupMockEnvironments()

	result := GetEnvironment("2", nil, nil)

	if result.ID != environment2.ID {
		t.Error("Expected environment 2, got", result)
	}

}
