package bodies_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBodies(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bodies Suite")
}
