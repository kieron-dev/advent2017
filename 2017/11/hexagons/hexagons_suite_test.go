package hexagons_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHexagons(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hexagons Suite")
}
