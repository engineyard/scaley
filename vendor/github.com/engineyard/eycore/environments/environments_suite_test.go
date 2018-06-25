package environments_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	httpmock "gopkg.in/jarcoal/httpmock.v1"

	"testing"
)

func TestEnvironments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Environments Suite")
}

var _ = BeforeSuite(func() {
	// block all HTTP requests
	httpmock.Activate()
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})
