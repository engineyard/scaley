package client_test

import (
	"fmt"
	"net/url"

	. "github.com/engineyard/eycore/client"
	"github.com/engineyard/eycore/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

type sausages struct {
	Sausages string
}

func (body *sausages) Body() []byte {
	return []byte(fmt.Sprintf(`{"sausages":"%s"}`, body.Sausages))
}

var _ = Describe("CoreApi", func() {
	var coreAPI *CoreAPI

	BeforeEach(func() {
		coreAPI = NewCoreAPI("api.example.com", "supersecret")
	})

	It("is a Client", func() {
		var i core.Client
		i = coreAPI
		Expect(i).To(Equal(coreAPI))
	})

	It("has a default Host", func() {
		defaultCoreAPI := NewCoreAPI("", "supersecret")
		baseurl := defaultCoreAPI.BaseURL
		Expect(baseurl.Host).To(Equal("api.engineyard.com"))
	})

	It("uses https", func() {
		baseurl := coreAPI.BaseURL
		Expect(baseurl.Scheme).To(Equal("https"))
	})

	Describe("Get", func() {
		var result []byte
		var apiError error

		Context("without Params", func() {
			var withoutParams string

			BeforeEach(func() {
				withoutParams = `{"params":{}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("GET", "https://api.example.com/without-params",
						httpmock.NewStringResponder(200, withoutParams))

					result, apiError = coreAPI.Get("/without-params", nil)
				})

				It("returns the proper content for the API GET", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withoutParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("GET", "https://api.example.com/without-params",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Get("/without-params", nil)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})

		})

		Context("with Params", func() {
			var withParams string
			var params url.Values

			BeforeEach(func() {
				params = url.Values{}
				params.Set("v", "1")

				withParams = `{"params":{"v":"1"}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("GET", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(200, withParams))

					result, apiError = coreAPI.Get("/with-params", params)
				})

				It("returns the proper content for the API GET", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("GET", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Get("/with-params", params)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})
		})
	})

	Describe("Post", func() {
		var result []byte
		var apiError error
		var data *sausages

		BeforeEach(func() {
			data = &sausages{"gold"}
		})

		Context("without Params", func() {
			var withoutParams string

			BeforeEach(func() {
				withoutParams = `{"params":{}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("POST", "https://api.example.com/without-params",
						httpmock.NewStringResponder(200, withoutParams))

					result, apiError = coreAPI.Post("/without-params", nil, data)
				})

				It("returns the proper content for the API POST", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withoutParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("POST", "https://api.example.com/without-params",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Post("/without-params", nil, data)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})

		})

		Context("with Params", func() {
			var withParams string
			var params url.Values

			BeforeEach(func() {
				params = url.Values{}
				params.Set("v", "1")

				withParams = `{"params":{"v":"1"}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("POST", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(200, withParams))

					result, apiError = coreAPI.Post("/with-params", params, data)
				})

				It("returns the proper content for the API POST", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("POST", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Post("/with-params", params, data)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})
		})
	})

	Describe("Put", func() {
		var result []byte
		var apiError error
		var data *sausages

		BeforeEach(func() {
			data = &sausages{"gold"}
		})

		Context("without Params", func() {
			var withoutParams string

			BeforeEach(func() {
				withoutParams = `{"params":{}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("PUT", "https://api.example.com/without-params",
						httpmock.NewStringResponder(200, withoutParams))

					result, apiError = coreAPI.Put("/without-params", nil, data)
				})

				It("returns the proper content for the API PUT", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withoutParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("PUT", "https://api.example.com/without-params",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Put("/without-params", nil, data)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})

		})

		Context("with Params", func() {
			var withParams string
			var params url.Values

			BeforeEach(func() {
				params = url.Values{}
				params.Set("v", "1")

				withParams = `{"params":{"v":"1"}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("PUT", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(200, withParams))

					result, apiError = coreAPI.Put("/with-params", params, data)
				})

				It("returns the proper content for the API PUT", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("PUT", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Put("/with-params", params, data)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})
		})
	})

	Describe("Delete", func() {
		var result []byte
		var apiError error

		Context("without Params", func() {
			var withoutParams string

			BeforeEach(func() {
				withoutParams = `{"params":{}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("DELETE", "https://api.example.com/without-params",
						httpmock.NewStringResponder(200, withoutParams))

					result, apiError = coreAPI.Delete("/without-params", nil)
				})

				It("returns the proper content for the API DELETE", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withoutParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("DELETE", "https://api.example.com/without-params",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Delete("/without-params", nil)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})

		})

		Context("with Params", func() {
			var withParams string
			var params url.Values

			BeforeEach(func() {
				params = url.Values{}
				params.Set("v", "1")

				withParams = `{"params":{"v":"1"}}`
			})

			Context("when there are no API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("DELETE", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(200, withParams))

					result, apiError = coreAPI.Delete("/with-params", params)
				})

				It("returns the proper content for the API DELETE", func() {
					var i []byte
					i = result
					Expect(i).To(Equal(result))
					Expect(string(result)).To(Equal(withParams))
				})

				It("returns a nil error", func() {
					Expect(apiError).To(BeNil())
				})
			})

			Context("when there are API errors", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder("DELETE", "https://api.example.com/with-params?v=1",
						httpmock.NewStringResponder(404, `{"errors":["onoes"]}`))

					result, apiError = coreAPI.Delete("/with-params", params)
				})

				It("returns a nil result", func() {
					Expect(result).To(BeNil())
				})

				It("returns a non-nil error", func() {
					Expect(apiError).NotTo(BeNil())
				})
			})
		})
	})
})
