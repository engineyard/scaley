package client_test

import (
	"fmt"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	. "github.com/engineyard/eycore/client"
)

type dumbBody struct{}

func (body *dumbBody) Body() []byte {
	return make([]byte, 0)
}

func WithAMockedAPI(doit func()) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://api.engineyard.com/users/current",
		httpmock.NewStringResponder(200,
			`{"user": {"id": "1", "name": "Current User", "email": "current@example.com"}}`))

	httpmock.RegisterResponder("POST", "https://api.engineyard.com/environments",
		httpmock.NewStringResponder(200,
			`{"environment": {"id":1, "name":"Environment 1"}}`))

	httpmock.RegisterResponder("PUT", "https://api.engineyard.com/environments/1",
		httpmock.NewStringResponder(200,
			`{"environment": {"id":1, "name":"Environment 1", "updated_at": "2017-01-01T00:00:00Z"}}`))

	httpmock.RegisterResponder("DELETE", "https://api.engineyard.com/environments/1",
		httpmock.NewStringResponder(200,
			`{"environment": {"id":1, "name":"Environment 1", "deleted_at": "2017-01-01T00:00:00Z"}}`))

	httpmock.RegisterResponder("DELETE", "https://api.engineyard.com/environments/2",
		httpmock.NewStringResponder(404, `{"errors":["Some error message"]}`))

	httpmock.RegisterResponder("DELETE", "https://api.engineyard.com/environments/3",
		httpmock.NewStringResponder(404, `{"errors":what?}`))

	httpmock.RegisterResponder("DELETE", "https://api.engineyard.com/environments/4",
		httpmock.NewStringResponder(500, "{}"))

	doit()

	httpmock.DeactivateAndReset()
}

func ExampleNewCoreAPI() {
	defaultClient := NewCoreAPI("", "")

	fmt.Println(defaultClient.BaseURL.String())

	api := NewCoreAPI("api.someotherhost.com", "")

	fmt.Println(api.BaseURL.String())
	// Output:
	// https://api.engineyard.com
	// https://api.someotherhost.com
}

func ExampleCoreAPI_Get() {
	WithAMockedAPI(func() {
		api := NewCoreAPI("", "supersecret")

		if result, err := api.Get("/users/current", nil); err == nil {
			fmt.Println(string(result))
		}
	})

	// Output:
	// {"user": {"id": "1", "name": "Current User", "email": "current@example.com"}}
}

func ExampleCoreAPI_Post() {
	WithAMockedAPI(func() {
		api := NewCoreAPI("", "supersecret")
		data := &dumbBody{}

		if result, err := api.Post("/environments", nil, data); err == nil {
			fmt.Println(string(result))
		}
	})

	// Output:
	// {"environment": {"id":1, "name":"Environment 1"}}
}

func ExampleCoreAPI_Put() {
	WithAMockedAPI(func() {
		api := NewCoreAPI("", "supersecret")
		data := &dumbBody{}

		if result, err := api.Put("/environments/1", nil, data); err == nil {
			fmt.Println(string(result))
		}
	})

	// Output:
	// {"environment": {"id":1, "name":"Environment 1", "updated_at": "2017-01-01T00:00:00Z"}}
}

func ExampleCoreAPI_Delete() {
	WithAMockedAPI(func() {
		api := NewCoreAPI("", "supersecret")

		if result, err := api.Delete("/environments/1", nil); err == nil {
			fmt.Println(string(result))
		}
	})

	// Output:
	// {"environment": {"id":1, "name":"Environment 1", "deleted_at": "2017-01-01T00:00:00Z"}}
}
