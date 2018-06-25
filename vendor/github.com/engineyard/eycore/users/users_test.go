package users_test

import (
	httpmock "gopkg.in/jarcoal/httpmock.v1"
	"net/url"

	"github.com/engineyard/eycore/client"

	. "github.com/engineyard/eycore/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users", func() {
	var api *client.CoreAPI

	BeforeEach(func() {
		api = client.NewCoreAPI("", "")
	})

	Describe("All", func() {
		var result []*Model

		Context("with multiple pages of results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[{"id":"1","name":"joe 1","email":"joe1@example.com"},{"id":"2","name":"joe 2","email":"joe2@example.com"},{"id":"3","name":"joe 3","email":"joe3@example.com"},{"id":"4","name":"joe 4","email":"joe4@example.com"},{"id":"5","name":"joe 5","email":"joe5@example.com"},{"id":"6","name":"joe 6","email":"joe6@example.com"},{"id":"7","name":"joe 7","email":"joe7@example.com"},{"id":"8","name":"joe 8","email":"joe8@example.com"},{"id":"9","name":"joe 9","email":"joe9@example.com"},{"id":"10","name":"joe 10","email":"joe10@example.com"},{"id":"11","name":"joe 11","email":"joe11@example.com"},{"id":"12","name":"joe 12","email":"joe12@example.com"},{"id":"13","name":"joe 13","email":"joe13@example.com"},{"id":"14","name":"joe 14","email":"joe14@example.com"},{"id":"15","name":"joe 15","email":"joe15@example.com"},{"id":"16","name":"joe 16","email":"joe16@example.com"},{"id":"17","name":"joe 17","email":"joe17@example.com"},{"id":"18","name":"joe 18","email":"joe18@example.com"},{"id":"19","name":"joe 19","email":"joe19@example.com"},{"id":"20","name":"joe 20","email":"joe20@example.com"},{"id":"21","name":"joe 21","email":"joe21@example.com"},{"id":"22","name":"joe 22","email":"joe22@example.com"},{"id":"23","name":"joe 23","email":"joe23@example.com"},{"id":"24","name":"joe 24","email":"joe24@example.com"},{"id":"25","name":"joe 25","email":"joe25@example.com"},{"id":"26","name":"joe 26","email":"joe26@example.com"},{"id":"27","name":"joe 27","email":"joe27@example.com"},{"id":"28","name":"joe 28","email":"joe28@example.com"},{"id":"29","name":"joe 29","email":"joe29@example.com"},{"id":"30","name":"joe 30","email":"joe30@example.com"},{"id":"31","name":"joe 31","email":"joe31@example.com"},{"id":"32","name":"joe 32","email":"joe32@example.com"},{"id":"33","name":"joe 33","email":"joe33@example.com"},{"id":"34","name":"joe 34","email":"joe34@example.com"},{"id":"35","name":"joe 35","email":"joe35@example.com"},{"id":"36","name":"joe 36","email":"joe36@example.com"},{"id":"37","name":"joe 37","email":"joe37@example.com"},{"id":"38","name":"joe 38","email":"joe38@example.com"},{"id":"39","name":"joe 39","email":"joe39@example.com"},{"id":"40","name":"joe 40","email":"joe40@example.com"},{"id":"41","name":"joe 41","email":"joe41@example.com"},{"id":"42","name":"joe 42","email":"joe42@example.com"},{"id":"43","name":"joe 43","email":"joe43@example.com"},{"id":"44","name":"joe 44","email":"joe44@example.com"},{"id":"45","name":"joe 45","email":"joe45@example.com"},{"id":"46","name":"joe 46","email":"joe46@example.com"},{"id":"47","name":"joe 47","email":"joe47@example.com"},{"id":"48","name":"joe 48","email":"joe48@example.com"},{"id":"49","name":"joe 49","email":"joe49@example.com"},{"id":"50","name":"joe 50","email":"joe50@example.com"},{"id":"51","name":"joe 51","email":"joe51@example.com"},{"id":"52","name":"joe 52","email":"joe52@example.com"},{"id":"53","name":"joe 53","email":"joe53@example.com"},{"id":"54","name":"joe 54","email":"joe54@example.com"},{"id":"55","name":"joe 55","email":"joe55@example.com"},{"id":"56","name":"joe 56","email":"joe56@example.com"},{"id":"57","name":"joe 57","email":"joe57@example.com"},{"id":"58","name":"joe 58","email":"joe58@example.com"},{"id":"59","name":"joe 59","email":"joe59@example.com"},{"id":"60","name":"joe 60","email":"joe60@example.com"},{"id":"61","name":"joe 61","email":"joe61@example.com"},{"id":"62","name":"joe 62","email":"joe62@example.com"},{"id":"63","name":"joe 63","email":"joe63@example.com"},{"id":"64","name":"joe 64","email":"joe64@example.com"},{"id":"65","name":"joe 65","email":"joe65@example.com"},{"id":"66","name":"joe 66","email":"joe66@example.com"},{"id":"67","name":"joe 67","email":"joe67@example.com"},{"id":"68","name":"joe 68","email":"joe68@example.com"},{"id":"69","name":"joe 69","email":"joe69@example.com"},{"id":"70","name":"joe 70","email":"joe70@example.com"},{"id":"71","name":"joe 71","email":"joe71@example.com"},{"id":"72","name":"joe 72","email":"joe72@example.com"},{"id":"73","name":"joe 73","email":"joe73@example.com"},{"id":"74","name":"joe 74","email":"joe74@example.com"},{"id":"75","name":"joe 75","email":"joe75@example.com"},{"id":"76","name":"joe 76","email":"joe76@example.com"},{"id":"77","name":"joe 77","email":"joe77@example.com"},{"id":"78","name":"joe 78","email":"joe78@example.com"},{"id":"79","name":"joe 79","email":"joe79@example.com"},{"id":"80","name":"joe 80","email":"joe80@example.com"},{"id":"81","name":"joe 81","email":"joe81@example.com"},{"id":"82","name":"joe 82","email":"joe82@example.com"},{"id":"83","name":"joe 83","email":"joe83@example.com"},{"id":"84","name":"joe 84","email":"joe84@example.com"},{"id":"85","name":"joe 85","email":"joe85@example.com"},{"id":"86","name":"joe 86","email":"joe86@example.com"},{"id":"87","name":"joe 87","email":"joe87@example.com"},{"id":"88","name":"joe 88","email":"joe88@example.com"},{"id":"89","name":"joe 89","email":"joe89@example.com"},{"id":"90","name":"joe 90","email":"joe90@example.com"},{"id":"91","name":"joe 91","email":"joe91@example.com"},{"id":"92","name":"joe 92","email":"joe92@example.com"},{"id":"93","name":"joe 93","email":"joe93@example.com"},{"id":"94","name":"joe 94","email":"joe94@example.com"},{"id":"95","name":"joe 95","email":"joe95@example.com"},{"id":"96","name":"joe 96","email":"joe96@example.com"},{"id":"97","name":"joe 97","email":"joe97@example.com"},{"id":"98","name":"joe 98","email":"joe98@example.com"},{"id":"99","name":"joe 99","email":"joe99@example.com"},{"id":"100","name":"joe 100","email":"joe100@example.com"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[{"id":"2","name":"jim","email":"jim@example.com"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?page=3&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[]}`))

				result = All(api, nil)
			})

			It("gets all results", func() {
				Expect(len(result)).To(Equal(101))
			})
		})

		Context("with no results", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[]}`))

				result = All(api, nil)
			})

			It("returns an empty array", func() {
				Expect(len(result)).To(Equal(0))
			})
		})

		Context("with params", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?name=foo&page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[{"id":"5","name":"bob","email":"bob@example.com"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/users?name=foo&page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"users":[]}`))

				params := url.Values{}
				params.Set("name", "foo")
				result = All(api, params)
			})

			It("passes the params along to the API", func() {
				Expect(result[0].Name).To(Equal("bob"))
			})
		})
	})

	Describe("Current", func() {
		var result *Model
		var err error

		Context("when there are no API errors", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET", "https://api.engineyard.com/users/current",
					httpmock.NewStringResponder(200,
						`{"user":{"id":"12345","name":"Joe User","email":"joe@example.com"}}`))

				result, err = Current(api)
			})

			It("returns the current user", func() {
				Expect(result.ID).To(Equal("12345"))
				Expect(result.Name).To(Equal("Joe User"))
				Expect(result.Email).To(Equal("joe@example.com"))
			})

			It("returns a nil error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when there are API errors", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder("GET", "https://api.engineyard.com/users/current",
					httpmock.NewStringResponder(404,
						`{"errors":["record not found"]}`))

				result, err = Current(api)
			})

			It("returns a nil user", func() {
				Expect(result).To(BeNil())
			})

			It("returns a non-nil error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})

})
