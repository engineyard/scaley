package flavors_test

import (
	httpmock "gopkg.in/jarcoal/httpmock.v1"

	"github.com/engineyard/eycore/accounts"
	"github.com/engineyard/eycore/client"

	. "github.com/engineyard/eycore/flavors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flavors", func() {
	var api *client.CoreAPI

	BeforeEach(func() {
		api = client.NewCoreAPI("", "")
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
					"https://api.engineyard.com/flavors?account=1&page=1&per_page=100",
					httpmock.NewStringResponder(200,
						`{"flavors":[{"id":"1"},{"id":"2"},{"id":"3"},{"id":"4"},{"id":"5"},{"id":"6"},{"id":"7"},{"id":"8"},{"id":"9"},{"id":"10"},{"id":"11"},{"id":"12"},{"id":"13"},{"id":"14"},{"id":"15"},{"id":"16"},{"id":"17"},{"id":"18"},{"id":"19"},{"id":"20"},{"id":"21"},{"id":"22"},{"id":"23"},{"id":"24"},{"id":"25"},{"id":"26"},{"id":"27"},{"id":"28"},{"id":"29"},{"id":"30"},{"id":"31"},{"id":"32"},{"id":"33"},{"id":"34"},{"id":"35"},{"id":"36"},{"id":"37"},{"id":"38"},{"id":"39"},{"id":"40"},{"id":"41"},{"id":"42"},{"id":"43"},{"id":"44"},{"id":"45"},{"id":"46"},{"id":"47"},{"id":"48"},{"id":"49"},{"id":"50"},{"id":"51"},{"id":"52"},{"id":"53"},{"id":"54"},{"id":"55"},{"id":"56"},{"id":"57"},{"id":"58"},{"id":"59"},{"id":"60"},{"id":"61"},{"id":"62"},{"id":"63"},{"id":"64"},{"id":"65"},{"id":"66"},{"id":"67"},{"id":"68"},{"id":"69"},{"id":"70"},{"id":"71"},{"id":"72"},{"id":"73"},{"id":"74"},{"id":"75"},{"id":"76"},{"id":"77"},{"id":"78"},{"id":"79"},{"id":"80"},{"id":"81"},{"id":"82"},{"id":"83"},{"id":"84"},{"id":"85"},{"id":"86"},{"id":"87"},{"id":"88"},{"id":"89"},{"id":"90"},{"id":"91"},{"id":"92"},{"id":"93"},{"id":"94"},{"id":"95"},{"id":"96"},{"id":"97"},{"id":"98"},{"id":"99"},{"id":"100"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/flavors?account=1&page=2&per_page=100",
					httpmock.NewStringResponder(200,
						`{"flavors":[{"id":"101"}]}`))

				httpmock.RegisterResponder("GET",
					"https://api.engineyard.com/flavors?account=1&page=3&per_page=100",
					httpmock.NewStringResponder(200,
						`{"flavors":[]}`))

				result = AllForAccount(api, parent, nil)
			})

			It("gets all results", func() {
				Expect(len(result)).To(Equal(101))
			})
		})

		Context("with no results", func() {})

		Context("with params", func() {})
	})

})
