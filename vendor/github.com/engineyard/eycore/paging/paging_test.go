package paging_test

import (
	"errors"
	"strconv"

	. "github.com/engineyard/eycore/paging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Paging", func() {
	Describe("MaxResults", func() {
		It("is the int representation of PerPage", func() {
			expected, _ := strconv.Atoi(PerPage)

			Expect(MaxResults()).To(Equal(expected))
		})
	})

	Describe("Page", func() {
		var err error

		Context("when given a nil error", func() {
			BeforeEach(func() {
				err = nil
			})

			It("is nil", func() {
				Expect(Page(err)).To(BeNil())
			})
		})

		Context("when given a FinalPage error", func() {
			BeforeEach(func() {
				err = errors.New(FinalPage)
			})

			It("is nil", func() {
				Expect(Page(err)).To(BeNil())
			})
		})

		Context("when given a non-nil, non-FinalPage error", func() {
			BeforeEach(func() {
				err = errors.New("this is a problem")
			})

			It("is the given error", func() {
				Expect(Page(err)).To(Equal(err))
			})
		})
	})

})
