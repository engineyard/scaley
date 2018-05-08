package paging_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPaging(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paging Suite")
}
