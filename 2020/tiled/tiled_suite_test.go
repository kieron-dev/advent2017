package tiled_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTiled(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tiled Suite")
}
