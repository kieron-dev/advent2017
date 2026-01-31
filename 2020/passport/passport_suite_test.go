package passport_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPassport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passport Suite")
}
