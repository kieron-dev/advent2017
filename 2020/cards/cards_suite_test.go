package cards_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cards Suite")
}
