package core_test

import (
	. "github.com/engineyard/eycore/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {

	Describe("NewError", func() {
		It("is a Error with the provided error string", func() {
			estring := "my sausages turned to mold :/"

			err := NewError(estring)

			Expect(err.ErrorString).To(Equal(estring))
		})
	})

})
