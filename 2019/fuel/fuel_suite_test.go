package fuel_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFuel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fuel Suite")
}
