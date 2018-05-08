package client_test

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/engineyard/eycore/accounts"
	. "github.com/engineyard/eycore/client"
	"github.com/engineyard/eycore/client/mockdata"
	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"
	"github.com/engineyard/eycore/requests"
	"github.com/engineyard/eycore/servers"
	"github.com/engineyard/eycore/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type wrapperBody struct {
	wrapped string
}

func (body *wrapperBody) Body() []byte {
	return []byte(body.wrapped)
}

var _ = Describe("MockAPI", func() {
	var mockAPI *MockAPI

	BeforeEach(func() {
		// Gotta seed dat mock dataset
		raw := `{"current_user":{"id":"1","name":"Kubey User","email":"kubey@example.com"},"users":[{"id":"1","name":"Kubey User","email":"kubey@example.com"},{"id":"2","name":"Bob User","email":"bob@example.com"}],"accounts":[{"id":"1","user_id":"1","name":"acct-1"},{"id":"2","user_id":"1","name":"acct-2"}],"environments":[{"id":1,"account_id":"1","name":"kubey-1"},{"id":2,"account_id":"2","name":"kubey-2"}],"servers":[{"id":1,"environment_id":1,"role":"app_master","provisioned_id":"i-0ebaabe0f4a30bfd5fcbae1f56bb65a5","flavor":{"id":"vanilla"},"public_hostname":"mclaughlinullrich.io","state":"running"},{"id":10,"environment_id":2,"role":"app","provisioned_id":"i-e1c6043fc55a87f8ec86016dd3127bae","flavor":{"id":"vanilla"},"public_hostname":"ziemann.info","state":"running"}],"requests":[{"id":"12345","type":"test_event"}]}`

		mockdata.Seed([]byte(raw))

		mockAPI = NewMockAPI("api.example.com", "supersecret")

	})

	It("is a Client", func() {
		var i core.Client
		i = mockAPI
		Expect(i).To(Equal(mockAPI))
	})

	It("has a default Host", func() {
		defaultMockAPI := NewMockAPI("", "supersecret")
		baseurl := defaultMockAPI.BaseURL
		Expect(baseurl.Host).To(Equal("api.engineyard.com"))
	})

	It("uses https", func() {
		baseurl := mockAPI.BaseURL
		Expect(baseurl.Scheme).To(Equal("https"))
	})

	Describe("Get", func() {
		var result []byte
		var apiError error

		Context("for users", func() {
			var wrapper struct {
				User *users.Model `json:"user"`
			}

			BeforeEach(func() {
				wrapper = struct {
					User *users.Model `json:"user"`
				}{nil}
			})

			Context("getting a collection", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/users", nil)
				})

				It("returns a collection of users as JSON", func() {
					uc := &users.Collection{}
					json.Unmarshal(result, uc)
					Expect(len(uc.Collected)).To(Equal(2))
				})
			})

			Context("getting a single user", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/users/2", nil)
				})

				It("is the requested user as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.User.ID).To(Equal("2"))
				})
			})

			Context("getting the current user", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/users/current", nil)
				})

				It("is the current user as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.User.ID).To(Equal("1"))
				})
			})

			Context("with params", func() {
				var uc *users.Collection
				var params url.Values

				BeforeEach(func() {
					params = url.Values{}
					params.Set("name", "Bob User")

					result, apiError = mockAPI.Get("/users", params)
					uc = &users.Collection{}
					json.Unmarshal(result, uc)
				})

				It("contains users that match the passed params", func() {
					Expect(len(uc.Collected)).To(Equal(1))
					Expect(uc.Collected[0].Name).To(Equal("Bob User"))
				})

				It("excludes non-matching users", func() {
					Expect(len(uc.Collected)).To(Equal(1))
					Expect(uc.Collected[0].Name).NotTo(Equal("Kubey User"))
				})

			})

			Context("with pages", func() {
				It("contains results for the first page", func() {
					params := url.Values{}
					params.Set("page", "1")

					collection := &users.Collection{}
					result, apiError = mockAPI.Get("/users", params)

					json.Unmarshal(result, collection)

					Expect(len(collection.Collected) > 0).To(BeTrue())
				})

				It("contains an empty array for any other page", func() {
					for page := 2; page <= 100; page++ {
						params := url.Values{}
						params.Set("page", strconv.Itoa(page))

						collection := &users.Collection{}
						result, apiError = mockAPI.Get("/users", params)

						json.Unmarshal(result, collection)

						Expect(len(collection.Collected)).To(Equal(0))
					}
				})
			})
		})

		Context("for accounts", func() {
			var wrapper struct {
				Account *accounts.Model `json:"account"`
			}

			BeforeEach(func() {
				wrapper = struct {
					Account *accounts.Model `json:"account"`
				}{nil}
			})

			Context("getting a collection", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/accounts", nil)
				})

				It("returns a collection of accounts as JSON", func() {
					ac := &accounts.Collection{}
					json.Unmarshal(result, ac)
					Expect(len(ac.Collected)).To(Equal(2))
				})
			})

			Context("getting a collection with a user scope", func() {
				var ac *accounts.Collection

				BeforeEach(func() {
					result, apiError = mockAPI.Get("/users/1/accounts", nil)
					ac = &accounts.Collection{}
					json.Unmarshal(result, ac)
				})

				It("includes accounts associated with the user", func() {
					found := false

					for _, account := range ac.Collected {
						if account.UserID == "1" {
							found = true
						}
					}

					Expect(found).To(BeTrue())
				})

				It("excludes accounts not associated with the user", func() {
					found := false

					for _, account := range ac.Collected {
						if account.UserID == "2" {
							found = true
						}
					}

					Expect(found).NotTo(BeTrue())
				})
			})

			Context("getting a single account", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/accounts/2", nil)
				})

				It("is the requested account as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Account.ID).To(Equal("2"))
				})
			})

			Context("getting a single account with a user scope", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/users/1/accounts/2", nil)
				})

				It("is the requested account as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Account.ID).To(Equal("2"))
				})
			})

			Context("with params", func() {
				var ac *accounts.Collection
				var params url.Values

				BeforeEach(func() {
					params = url.Values{}
					params.Set("name", "acct-2")

					result, apiError = mockAPI.Get("/accounts", params)
					ac = &accounts.Collection{}
					json.Unmarshal(result, ac)
				})

				It("contains accounts that match the passed params", func() {
					Expect(len(ac.Collected)).To(Equal(1))
					Expect(ac.Collected[0].Name).To(Equal("acct-2"))
				})

				It("excludes non-matching accounts", func() {
					Expect(len(ac.Collected)).To(Equal(1))
					Expect(ac.Collected[0].Name).NotTo(Equal("acct-1"))
				})

			})

			Context("with pages", func() {
				It("contains results for the first page", func() {
					params := url.Values{}
					params.Set("page", "1")

					collection := &accounts.Collection{}
					result, apiError = mockAPI.Get("/accounts", params)

					json.Unmarshal(result, collection)

					Expect(len(collection.Collected) > 0).To(BeTrue())
				})

				It("contains an empty array for any other page", func() {
					for page := 2; page <= 100; page++ {
						params := url.Values{}
						params.Set("page", strconv.Itoa(page))

						collection := &accounts.Collection{}
						result, apiError = mockAPI.Get("/accounts", params)

						json.Unmarshal(result, collection)

						Expect(len(collection.Collected)).To(Equal(0))
					}
				})
			})

		})

		Context("for environments", func() {
			var wrapper struct {
				Environment *environments.Model `json:"environment"`
			}

			BeforeEach(func() {
				wrapper = struct {
					Environment *environments.Model `json:"environment"`
				}{nil}
			})

			Context("getting a collection", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/environments", nil)
				})

				It("returns a collection of environments as JSON", func() {
					ec := &environments.Collection{}
					json.Unmarshal(result, ec)
					Expect(len(ec.Collected)).To(Equal(2))
				})
			})

			Context("getting a collection with an account scope", func() {
				var ec *environments.Collection

				BeforeEach(func() {
					result, apiError = mockAPI.Get("/accounts/1/environments", nil)
					ec = &environments.Collection{}
					json.Unmarshal(result, ec)
				})

				It("includes environments associated with the account", func() {
					found := false

					for _, environment := range ec.Collected {
						if environment.AccountID == "1" {
							found = true
						}
					}

					Expect(found).To(BeTrue())
				})

				It("excludes environments not associated with the account", func() {
					found := false

					for _, environment := range ec.Collected {
						if environment.AccountID == "2" {
							found = true
						}
					}

					Expect(found).NotTo(BeTrue())
				})
			})

			Context("getting a single environment", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/environments/2", nil)
				})

				It("is the requested environment as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Environment.ID).To(Equal(2))
				})
			})

			Context("getting a single environment with an account scope", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/accounts/2/environments/2", nil)
				})

				It("is the requested environment as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Environment.ID).To(Equal(2))
				})
			})

			Context("with params", func() {
				var ec *environments.Collection
				var params url.Values

				BeforeEach(func() {
					params = url.Values{}
					params.Set("name", "kubey-2")

					result, apiError = mockAPI.Get("/environments", params)
					ec = &environments.Collection{}
					json.Unmarshal(result, ec)
				})

				It("contains environments that match the passed params", func() {
					Expect(len(ec.Collected)).To(Equal(1))
					Expect(ec.Collected[0].Name).To(Equal("kubey-2"))
				})

				It("excludes non-matching environments", func() {
					Expect(len(ec.Collected)).To(Equal(1))
					Expect(ec.Collected[0].Name).NotTo(Equal("kubey-1"))
				})

			})

			Context("with pages", func() {
				It("contains results for the first page", func() {
					params := url.Values{}
					params.Set("page", "1")

					collection := &environments.Collection{}
					result, apiError = mockAPI.Get("/environments", params)

					json.Unmarshal(result, collection)

					Expect(len(collection.Collected) > 0).To(BeTrue())
				})

				It("contains an empty array for any other page", func() {
					for page := 2; page <= 100; page++ {
						params := url.Values{}
						params.Set("page", strconv.Itoa(page))

						collection := &environments.Collection{}
						result, apiError = mockAPI.Get("/environments", params)

						json.Unmarshal(result, collection)

						Expect(len(collection.Collected)).To(Equal(0))
					}
				})
			})

		})

		Context("for servers", func() {
			var wrapper struct {
				Server *servers.Model `json:"server"`
			}

			BeforeEach(func() {
				wrapper = struct {
					Server *servers.Model `json:"server"`
				}{nil}
			})

			Context("getting a collection", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/servers", nil)
				})

				It("returns a collection of servers as JSON", func() {
					sc := &servers.Collection{}
					json.Unmarshal(result, sc)
					Expect(len(sc.Collected)).To(Equal(2))
				})
			})

			Context("getting a collection with an environment scope", func() {
				var sc *servers.Collection

				BeforeEach(func() {
					result, apiError = mockAPI.Get("/environments/1/servers", nil)
					sc = &servers.Collection{}
					json.Unmarshal(result, sc)
				})

				It("includes servers associated with the environment", func() {
					found := false

					for _, server := range sc.Collected {
						if server.EnvironmentID == 1 {
							found = true
						}
					}

					Expect(found).To(BeTrue())
				})

				It("excludes servers not associated with the environment", func() {
					found := false

					for _, server := range sc.Collected {
						if server.EnvironmentID == 2 {
							found = true
						}
					}

					Expect(found).NotTo(BeTrue())
				})
			})

			Context("getting a single server", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/servers/10", nil)
				})

				It("is the requested server as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Server.ID).To(Equal(10))
				})
			})

			Context("getting a single server with an environment scope", func() {
				BeforeEach(func() {
					result, apiError = mockAPI.Get("/environments/2/servers/10", nil)
				})

				It("is the requested server as JSON", func() {
					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Server.ID).To(Equal(10))
				})
			})

			Context("with params", func() {
				var sc *servers.Collection
				var params url.Values

				BeforeEach(func() {
					params = url.Values{}
					params.Set("role", "app_master")

					result, apiError = mockAPI.Get("/servers", params)
					sc = &servers.Collection{}
					json.Unmarshal(result, sc)
				})

				It("contains servers that match the passed params", func() {
					Expect(len(sc.Collected)).To(Equal(1))
					Expect(sc.Collected[0].Role).To(Equal("app_master"))
				})

				It("excludes non-matching servers", func() {
					Expect(len(sc.Collected)).To(Equal(1))
					Expect(sc.Collected[0].Role).NotTo(Equal("app"))
				})

			})

			Context("with pages", func() {
				It("contains results for the first page", func() {
					params := url.Values{}
					params.Set("page", "1")

					collection := &servers.Collection{}
					result, apiError = mockAPI.Get("/servers", params)

					json.Unmarshal(result, collection)

					Expect(len(collection.Collected) > 0).To(BeTrue())
				})

				It("contains an empty array for any other page", func() {
					for page := 2; page <= 100; page++ {
						params := url.Values{}
						params.Set("page", strconv.Itoa(page))

						collection := &servers.Collection{}
						result, apiError = mockAPI.Get("/servers", params)

						json.Unmarshal(result, collection)

						Expect(len(collection.Collected)).To(Equal(0))
					}
				})
			})

		})

		Context("for requests", func() {
			var stored string
			var wrapper struct {
				Request *requests.Model `json:"request"`
			}

			BeforeEach(func() {
				wrapper = struct {
					Request *requests.Model `json:"request"`
				}{nil}

				stored = mockdata.GetRequest("12345", nil, nil).FinishedAt

				result, apiError = mockAPI.Get("/requests/12345", nil)

				json.Unmarshal(result, &wrapper)
			})

			It("is a JSON representation of the request in question", func() {
				Expect(wrapper.Request.ID).To(Equal("12345"))
				Expect(wrapper.Request.Type).To(Equal("test_event"))
			})

			It("sets FinishedAt on the stored Request", func() {
				Expect(len(stored)).To(Equal(0))

				Expect(len(mockdata.GetRequest("12345", nil, nil).FinishedAt)).
					NotTo(Equal(0))
			})

		})
	})

	/*Describe("Post", func() {
		var result []byte
		var apiError error
		var data []byte

	})*/

	Describe("Delete", func() {
		var result []byte
		var apiError error

		Context("for environments", func() {
			Context("deleting an environment", func() {
				var reqCount int

				BeforeEach(func() {
					reqCount = len(mockdata.Requests())
					result, apiError = mockAPI.Delete("/environments/1", nil)
				})

				It("adds a destroy environment request to the dataset", func() {
					Expect(len(mockdata.Requests())).NotTo(Equal(reqCount))
				})

				It("returns a JSON representation of the added destroy environment request", func() {
					wrapper := struct {
						Request *requests.Model `json:"request"`
					}{nil}

					json.Unmarshal(result, &wrapper)

					Expect(wrapper.Request.Type).To(Equal("destroy_environment"))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})
		})
	})

	Describe("Post", func() {
		var data core.Body
		var baseResult []byte
		var scopedResult []byte
		var baseError error
		var scopedError error

		Context("for users", func() {
			// Client is not allowed to create users
			Context("adding a user", func() {
				BeforeEach(func() {
					data = &wrapperBody{`{"user":{"name":"Jim User", "email":"jim@example.com"}}`}
					baseResult, baseError = mockAPI.Post("/users", nil, data)
				})
				It("returns nil data", func() {
					Expect(baseResult).To(BeNil())
				})

				It("returns an error regarding client capabilities", func() {
					Expect(baseError).NotTo(BeNil())
					Expect(baseError.Error()).To(Equal(IllegalOperation))
				})
			})
		})

		Context("for accounts", func() {
			// Client is not allowed to create accounts
			Context("adding an account", func() {
				BeforeEach(func() {
					data = &wrapperBody{`{"account":{"name":"Jim Account"}}`}
					baseResult, baseError = mockAPI.Post("/accounts", nil, data)
					scopedResult, scopedError = mockAPI.Post("/users/1/accounts", nil, data)
				})

				It("returns nil data", func() {
					Expect(baseResult).To(BeNil())
					Expect(scopedResult).To(BeNil())
				})

				It("returns an error regarding client capabilities", func() {
					Expect(baseError.Error()).To(Equal(IllegalOperation))
					Expect(scopedError.Error()).To(Equal(IllegalOperation))
				})
			})
		})

		Context("for environments", func() {
			var envCount int

			Context("adding an environment", func() {
				BeforeEach(func() {
					envCount = len(mockdata.Environments())
					data = &wrapperBody{`{"environment":{"name":"Jim Environment","account_id":"1"}}`}
					baseResult, baseError = mockAPI.Post("/environments", nil, data)
					scopedResult, scopedError = mockAPI.Post("/accounts/1/environments", nil, data)
				})

				It("adds a new environment to the dataset", func() {
					Expect(len(mockdata.Environments())).NotTo(Equal(envCount))
				})

				It("returns a JSON representation of the new environment", func() {
					wrapper := struct {
						Environment *environments.Model `json:"environment"`
					}{nil}

					json.Unmarshal(baseResult, &wrapper)
					Expect(wrapper.Environment.Name).To(Equal("Jim Environment"))

					wrapper.Environment = nil
					json.Unmarshal(scopedResult, &wrapper)
					Expect(wrapper.Environment.Name).To(Equal("Jim Environment"))
				})

				It("returns a nil error", func() {
					Expect(baseError).To(BeNil())
					Expect(scopedError).To(BeNil())
				})
			})

			Context("booting an environment", func() {
				var reqCount int

				BeforeEach(func() {
					reqCount = len(mockdata.Requests())
					data = &wrapperBody{`{"cluster_configuration":{"configuration":{}}}`}
					baseResult, baseError = mockAPI.Post("/environments/1/boot", nil, data)
				})

				It("adds a boot request to the dataset", func() {
					Expect(len(mockdata.Requests())).NotTo(Equal(reqCount))
				})

				It("returns a JSON representation of the added boot request", func() {
					wrapper := struct {
						Request *requests.Model `json:"request"`
					}{nil}

					json.Unmarshal(baseResult, &wrapper)

					Expect(wrapper.Request.Type).To(Equal("start_environment"))
				})

				It("returns a nil error", func() {
					Expect(baseError).To(BeNil())
				})
			})
		})

		Context("for servers", func() {
			var reqCount int

			Context("adding a server", func() {
				BeforeEach(func() {
					reqCount = len(mockdata.Requests())
					data = &wrapperBody{`{"environment":1,"server":{"flavor":"m3.large","role":"util"}}`}
					baseResult, baseError = mockAPI.Post("/servers", nil, data)
					scopedResult, scopedError = mockAPI.Post("/environments/1/servers", nil, data)
				})

				It("adds a server creation request to the dataset", func() {
					Expect(len(mockdata.Requests()) - reqCount).To(Equal(2))
				})

				It("returns a JSON represenation of the added server create request", func() {
					wrapper := struct {
						Request *requests.Model `json:"request"`
					}{nil}

					json.Unmarshal(baseResult, &wrapper)
					Expect(wrapper.Request.Type).To(Equal("provision_server"))

					wrapper.Request = nil

					json.Unmarshal(scopedResult, &wrapper)
					Expect(wrapper.Request.Type).To(Equal("provision_server"))
				})

				It("returns a nil error", func() {
					Expect(baseError).To(BeNil())
				})
			})
		})

	})
})
