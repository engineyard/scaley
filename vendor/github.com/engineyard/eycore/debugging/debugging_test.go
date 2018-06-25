package debugging_test

import (
	"os"

	. "github.com/engineyard/eycore/debugging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Debugging", func() {
	AfterEach(func() {
		os.Unsetenv("DEBUG")
		os.Unsetenv("SHOWLOGS")
	})

	Describe("Enabled", func() {
		Context("when the DEBUG env is set", func() {
			It("is true", func() {
				os.Setenv("DEBUG", "1")
				Expect(Enabled()).To(Equal(true))
				os.Unsetenv("DEBUG")
			})
		})

		Context("when the DEBUG env is not set", func() {
			It("is false", func() {
				os.Unsetenv("DEBUG")
				Expect(Enabled()).To(Equal(false))

			})
		})
	})

	Describe("Live", func() {
		Context("when debugging is enabled", func() {
			BeforeEach(func() {
				os.Setenv("DEBUG", "1")
			})

			Context("when SHOWLOGS is set", func() {
				BeforeEach(func() {
					os.Setenv("SHOWLOGS", "1")
				})

				It("is true", func() {
					Expect(Live()).To(Equal(true))
				})
			})

			Context("when SHOWLOGS is not set", func() {
				BeforeEach(func() {
					os.Unsetenv("SHOWLOGS")
				})

				It("is false", func() {
					Expect(Live()).To(Equal(false))
				})
			})
		})

		Context("when debugging is not enabled", func() {
			BeforeEach(func() {
				os.Unsetenv("DEBUG")
			})

			It("is false", func() {
				Expect(Live()).To(Equal(false))
			})
		})
	})

})
