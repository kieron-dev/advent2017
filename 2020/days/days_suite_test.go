package days_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDays(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Days Suite")
}
