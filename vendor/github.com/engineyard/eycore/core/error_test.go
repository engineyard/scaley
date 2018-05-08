package core_test

import (
	. "github.com/engineyard/eycore/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {

	Describe("Error", func() {
		var err *Error

		Context("when no error string was provided", func() {
			BeforeEach(func() {
				err = &Error{}
			})

			It("is an empty string", func() {
				Expect(err.Error()).To(Equal(""))
			})
		})

		Context("when an error string was provided", func() {
			var estring string

			BeforeEach(func() {
				estring = "I am an error!"
				err = &Error{ErrorString: estring}
			})

			It("is the error string in question", func() {
				Expect(err.Error()).To(Equal(estring))
			})
		})
	})

})
