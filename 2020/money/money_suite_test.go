package money_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMoney(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Money Suite")
}
