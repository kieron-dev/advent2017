package registers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRegisters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registers Suite")
}
