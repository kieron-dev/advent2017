package tractor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTractor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tractor Suite")
}
