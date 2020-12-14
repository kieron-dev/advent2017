package masked_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMasked(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Masked Suite")
}
