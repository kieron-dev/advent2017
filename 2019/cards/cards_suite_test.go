package cards_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cards Suite")
}
