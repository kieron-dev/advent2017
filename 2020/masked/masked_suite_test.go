package masked_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMasked(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Masked Suite")
}
