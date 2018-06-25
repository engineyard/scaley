package mockdata

import (
	"net/url"
	"testing"

	"github.com/engineyard/eycore/users"
)

func TestGetUsersWithoutParams(t *testing.T) {
	user1 := &users.Model{
		ID:    "1",
		Name:  "Jim",
		Email: "jim@example.com",
	}

	user2 := &users.Model{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@example.com",
	}

	userStore = []*users.Model{user1, user2}

	result := GetUsers(nil, url.Values{})

	if len(result) != 2 {
		t.Error("Expected all users request to return all users, got", result)
	}
}

func TestGetUsersByName(t *testing.T) {
	user1 := &users.Model{
		ID:    "1",
		Name:  "Jim",
		Email: "jim@example.com",
	}

	user2 := &users.Model{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@example.com",
	}

	userStore = []*users.Model{user1, user2}

	params := url.Values{}
	params.Set("name", "Bob")

	result := GetUsers(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	user := result[0]
	if user.ID != user2.ID {
		t.Error("Expected the Bob user, got", result[0])
	}

}

func TestGetUsersByEmail(t *testing.T) {
	user1 := &users.Model{
		ID:    "1",
		Name:  "Jim",
		Email: "jim@example.com",
	}

	user2 := &users.Model{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@example.com",
	}

	userStore = []*users.Model{user1, user2}

	params := url.Values{}
	params.Set("email", "jim@example.com")

	result := GetUsers(nil, params)

	if len(result) > 1 {
		t.Error("Expected only one result, got", len(result))
	}

	user := result[0]
	if user.ID != user1.ID {
		t.Error("Expected the Jim user, got", result[0])
	}

}

func TestGetUser(t *testing.T) {
	user1 := &users.Model{
		ID:    "1",
		Name:  "Jim",
		Email: "jim@example.com",
	}

	user2 := &users.Model{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@example.com",
	}

	userStore = []*users.Model{user1, user2}

	result := GetUser("2", nil, nil)

	if result.ID != user2.ID {
		t.Error("Expected the Bob user, got", result)
	}

}

func TestGetUserCurrent(t *testing.T) {
	user1 := &users.Model{
		ID:    "1",
		Name:  "Jim",
		Email: "jim@example.com",
	}

	user2 := &users.Model{
		ID:    "2",
		Name:  "Bob",
		Email: "bob@example.com",
	}

	userStore = []*users.Model{user1, user2}
	currentUser = &users.Model{
		ID:    "31337",
		Name:  "Me",
		Email: "me@example.com",
	}

	result := GetUser("current", nil, nil)

	if result.ID != currentUser.ID {
		t.Error("Expected the current user, got", result)
	}

}
