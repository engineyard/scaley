package eycore_test

import (
	"fmt"
	"reflect"

	. "github.com/engineyard/eycore"
	"github.com/engineyard/eycore/core"

	"github.com/ess/mockable"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Eycore", func() {
	Describe("NewCoreAPI", func() {
		var result core.Client

		AfterEach(func() {
			mockable.Disable()
		})

		Context("when mocking is not enabled", func() {
			BeforeEach(func() {
				mockable.Disable()

				result = NewClient("", "")
			})

			It("returns a CoreAPI", func() {
				Expect(fmt.Sprintf("%s", reflect.TypeOf(result))).
					To(Equal("*client.CoreAPI"))
			})
		})

		Context("when mocking is enabled", func() {
			BeforeEach(func() {
				mockable.Enable()

				result = NewClient("", "")
			})

			It("returns a MockAPI", func() {
				Expect(fmt.Sprintf("%s", reflect.TypeOf(result))).
					To(Equal("*client.MockAPI"))
			})
		})
	})
})
