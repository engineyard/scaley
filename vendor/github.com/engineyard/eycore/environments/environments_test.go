package environments_test

import (
	httpmock "gopkg.in/jarcoal/httpmock.v1"
	"net/url"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/client"

	. "github.com/engineyard/eycore/environments"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Environments", func() {
	var api *client.CoreAPI

	BeforeEach(func() {
		api = client.NewCoreAPI("", "")
	})

	Describe("All", func() {
		var result []*Model

		Context("with multiple pages of results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":1,"name":"joe 1"},{"id":2,"name":"joe 2"},{"id":3,"name":"joe 3"},{"id":4,"name":"joe 4"},{"id":5,"name":"joe 5"},{"id":6,"name":"joe 6"},{"id":7,"name":"joe 7"},{"id":8,"name":"joe 8"},{"id":9,"name":"joe 9"},{"id":10,"name":"joe 10"},{"id":11,"name":"joe 11"},{"id":12,"name":"joe 12"},{"id":13,"name":"joe 13"},{"id":14,"name":"joe 14"},{"id":15,"name":"joe 15"},{"id":16,"name":"joe 16"},{"id":17,"name":"joe 17"},{"id":18,"name":"joe 18"},{"id":19,"name":"joe 19"},{"id":20,"name":"joe 20"},{"id":21,"name":"joe 21"},{"id":22,"name":"joe 22"},{"id":23,"name":"joe 23"},{"id":24,"name":"joe 24"},{"id":25,"name":"joe 25"},{"id":26,"name":"joe 26"},{"id":27,"name":"joe 27"},{"id":28,"name":"joe 28"},{"id":29,"name":"joe 29"},{"id":30,"name":"joe 30"},{"id":31,"name":"joe 31"},{"id":32,"name":"joe 32"},{"id":33,"name":"joe 33"},{"id":34,"name":"joe 34"},{"id":35,"name":"joe 35"},{"id":36,"name":"joe 36"},{"id":37,"name":"joe 37"},{"id":38,"name":"joe 38"},{"id":39,"name":"joe 39"},{"id":40,"name":"joe 40"},{"id":41,"name":"joe 41"},{"id":42,"name":"joe 42"},{"id":43,"name":"joe 43"},{"id":44,"name":"joe 44"},{"id":45,"name":"joe 45"},{"id":46,"name":"joe 46"},{"id":47,"name":"joe 47"},{"id":48,"name":"joe 48"},{"id":49,"name":"joe 49"},{"id":50,"name":"joe 50"},{"id":51,"name":"joe 51"},{"id":52,"name":"joe 52"},{"id":53,"name":"joe 53"},{"id":54,"name":"joe 54"},{"id":55,"name":"joe 55"},{"id":56,"name":"joe 56"},{"id":57,"name":"joe 57"},{"id":58,"name":"joe 58"},{"id":59,"name":"joe 59"},{"id":60,"name":"joe 60"},{"id":61,"name":"joe 61"},{"id":62,"name":"joe 62"},{"id":63,"name":"joe 63"},{"id":64,"name":"joe 64"},{"id":65,"name":"joe 65"},{"id":66,"name":"joe 66"},{"id":67,"name":"joe 67"},{"id":68,"name":"joe 68"},{"id":69,"name":"joe 69"},{"id":70,"name":"joe 70"},{"id":71,"name":"joe 71"},{"id":72,"name":"joe 72"},{"id":73,"name":"joe 73"},{"id":74,"name":"joe 74"},{"id":75,"name":"joe 75"},{"id":76,"name":"joe 76"},{"id":77,"name":"joe 77"},{"id":78,"name":"joe 78"},{"id":79,"name":"joe 79"},{"id":80,"name":"joe 80"},{"id":81,"name":"joe 81"},{"id":82,"name":"joe 82"},{"id":83,"name":"joe 83"},{"id":84,"name":"joe 84"},{"id":85,"name":"joe 85"},{"id":86,"name":"joe 86"},{"id":87,"name":"joe 87"},{"id":88,"name":"joe 88"},{"id":89,"name":"joe 89"},{"id":90,"name":"joe 90"},{"id":91,"name":"joe 91"},{"id":92,"name":"joe 92"},{"id":93,"name":"joe 93"},{"id":94,"name":"joe 94"},{"id":95,"name":"joe 95"},{"id":96,"name":"joe 96"},{"id":97,"name":"joe 97"},{"id":98,"name":"joe 98"},{"id":99,"name":"joe 99"},{"id":100,"name":"joe 100"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":2,"name":"jim"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?page=3&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				result = All(api, nil)
			})

			It("gets all results", func() {
				Expect(len(result)).To(Equal(101))
			})
		})

		Context("with no results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				result = All(api, nil)
			})

			It("returns an empty array", func() {
				Expect(len(result)).To(Equal(0))
			})
		})

		Context("with params", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?name=foo&page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":5,"name":"bob"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/environments?name=foo&page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				params := url.Values{}
				params.Set("name", "foo")
				result = All(api, params)
			})

			It("passes the params along to the API", func() {
				Expect(result[0].Name).To(Equal("bob"))
			})
		})
	})

	Describe("AllForAccount", func() {
		var result []*Model
		var parent *accounts.Model

		BeforeEach(func() {
			parent = &accounts.Model{ID: "1"}
		})

		Context("with multiple pages of results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":1,"name":"joe 1"},{"id":2,"name":"joe 2"},{"id":3,"name":"joe 3"},{"id":4,"name":"joe 4"},{"id":5,"name":"joe 5"},{"id":6,"name":"joe 6"},{"id":7,"name":"joe 7"},{"id":8,"name":"joe 8"},{"id":9,"name":"joe 9"},{"id":10,"name":"joe 10"},{"id":11,"name":"joe 11"},{"id":12,"name":"joe 12"},{"id":13,"name":"joe 13"},{"id":14,"name":"joe 14"},{"id":15,"name":"joe 15"},{"id":16,"name":"joe 16"},{"id":17,"name":"joe 17"},{"id":18,"name":"joe 18"},{"id":19,"name":"joe 19"},{"id":20,"name":"joe 20"},{"id":21,"name":"joe 21"},{"id":22,"name":"joe 22"},{"id":23,"name":"joe 23"},{"id":24,"name":"joe 24"},{"id":25,"name":"joe 25"},{"id":26,"name":"joe 26"},{"id":27,"name":"joe 27"},{"id":28,"name":"joe 28"},{"id":29,"name":"joe 29"},{"id":30,"name":"joe 30"},{"id":31,"name":"joe 31"},{"id":32,"name":"joe 32"},{"id":33,"name":"joe 33"},{"id":34,"name":"joe 34"},{"id":35,"name":"joe 35"},{"id":36,"name":"joe 36"},{"id":37,"name":"joe 37"},{"id":38,"name":"joe 38"},{"id":39,"name":"joe 39"},{"id":40,"name":"joe 40"},{"id":41,"name":"joe 41"},{"id":42,"name":"joe 42"},{"id":43,"name":"joe 43"},{"id":44,"name":"joe 44"},{"id":45,"name":"joe 45"},{"id":46,"name":"joe 46"},{"id":47,"name":"joe 47"},{"id":48,"name":"joe 48"},{"id":49,"name":"joe 49"},{"id":50,"name":"joe 50"},{"id":51,"name":"joe 51"},{"id":52,"name":"joe 52"},{"id":53,"name":"joe 53"},{"id":54,"name":"joe 54"},{"id":55,"name":"joe 55"},{"id":56,"name":"joe 56"},{"id":57,"name":"joe 57"},{"id":58,"name":"joe 58"},{"id":59,"name":"joe 59"},{"id":60,"name":"joe 60"},{"id":61,"name":"joe 61"},{"id":62,"name":"joe 62"},{"id":63,"name":"joe 63"},{"id":64,"name":"joe 64"},{"id":65,"name":"joe 65"},{"id":66,"name":"joe 66"},{"id":67,"name":"joe 67"},{"id":68,"name":"joe 68"},{"id":69,"name":"joe 69"},{"id":70,"name":"joe 70"},{"id":71,"name":"joe 71"},{"id":72,"name":"joe 72"},{"id":73,"name":"joe 73"},{"id":74,"name":"joe 74"},{"id":75,"name":"joe 75"},{"id":76,"name":"joe 76"},{"id":77,"name":"joe 77"},{"id":78,"name":"joe 78"},{"id":79,"name":"joe 79"},{"id":80,"name":"joe 80"},{"id":81,"name":"joe 81"},{"id":82,"name":"joe 82"},{"id":83,"name":"joe 83"},{"id":84,"name":"joe 84"},{"id":85,"name":"joe 85"},{"id":86,"name":"joe 86"},{"id":87,"name":"joe 87"},{"id":88,"name":"joe 88"},{"id":89,"name":"joe 89"},{"id":90,"name":"joe 90"},{"id":91,"name":"joe 91"},{"id":92,"name":"joe 92"},{"id":93,"name":"joe 93"},{"id":94,"name":"joe 94"},{"id":95,"name":"joe 95"},{"id":96,"name":"joe 96"},{"id":97,"name":"joe 97"},{"id":98,"name":"joe 98"},{"id":99,"name":"joe 99"},{"id":100,"name":"joe 100"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":2,"name":"jim"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?page=3&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				result = AllForAccount(api, parent, nil)
			})

			It("gets all results", func() {
				Expect(len(result)).To(Equal(101))
			})
		})

		Context("with no results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				result = AllForAccount(api, parent, nil)
			})

			It("returns an empty array", func() {
				Expect(len(result)).To(Equal(0))
			})
		})

		Context("with params", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?name=foo&page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[{"id":5,"name":"bob"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/accounts/1/environments?name=foo&page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"environments":[]}`))

				params := url.Values{}
				params.Set("name", "foo")
				result = AllForAccount(api, parent, params)
			})

			It("passes the params along to the API", func() {
				Expect(result[0].Name).To(Equal("bob"))
			})
		})
	})

})
