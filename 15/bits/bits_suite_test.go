package bits_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBits(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bits Suite")
}
