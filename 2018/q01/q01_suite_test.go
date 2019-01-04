package q01_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQ01(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Q01 Suite")
}
