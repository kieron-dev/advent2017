package donut_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDonut(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Donut Suite")
}
