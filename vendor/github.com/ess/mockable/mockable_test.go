package mockable_test

import (
	"os"

	. "github.com/ess/mockable"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const mockableEnvVar = "MOCKABLE"

var _ = Describe("Mockable", func() {
	Describe("Mocked", func() {
		var result bool

		AfterEach(func() {
			Disable()
		})

		Context("when MOCKABLE is not set", func() {
			BeforeEach(func() {
				os.Unsetenv(mockableEnvVar)

				result = Mocked()
			})

			It("is false", func() {
				Expect(result).To(Equal(false))
			})
		})

		Context("when MOCKABLE is not set", func() {
			BeforeEach(func() {
				os.Setenv(mockableEnvVar, "anything")

				result = Mocked()
			})

			It("is true", func() {
				Expect(result).To(Equal(true))
			})
		})
	})

	Describe("Enable", func() {
		Context("when mocking is not enabled", func() {
			BeforeEach(func() {
				Disable()
			})

			It("turns mocking on", func() {
				Expect(Mocked()).To(Equal(false))

				Enable()

				Expect(Mocked()).To(Equal(true))
			})
		})

		Context("when mocking is enabled", func() {
			var value = "sure, why not"

			BeforeEach(func() {
				os.Setenv(mockableEnvVar, value)
			})

			It("has no apparent effect", func() {
				Enable()
				Expect(Mocked()).To(Equal(true))
				Expect(os.Getenv(mockableEnvVar)).To(Equal(value))
			})
		})
	})

	Describe("Disable", func() {
		Context("when mocking is enabled", func() {
			BeforeEach(func() {
				Enable()
			})

			It("disabled mocking", func() {
				Expect(Mocked()).To(Equal(true))
				Disable()
				Expect(Mocked()).To(Equal(false))
			})
		})

		Context("when mocking is not enabled", func() {
			It("has no apparent effect", func() {
				Expect(Mocked()).To(Equal(false))
				Disable()
				Expect(Mocked()).To(Equal(false))
			})
		})
	})
})
