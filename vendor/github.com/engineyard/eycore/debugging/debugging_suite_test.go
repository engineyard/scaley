package debugging_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDebugging(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Debugging Suite")
}
