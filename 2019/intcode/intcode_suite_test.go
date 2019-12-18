package intcode_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestComputer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intcode Suite")
}
