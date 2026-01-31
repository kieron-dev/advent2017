package cups_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCups(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cups Suite")
}
