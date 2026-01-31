package gameoflife_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGameoflife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gameoflife Suite")
}
