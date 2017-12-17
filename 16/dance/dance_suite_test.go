package dance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dance Suite")
}
