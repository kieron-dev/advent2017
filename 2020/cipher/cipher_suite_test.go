package cipher_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCipher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cipher Suite")
}
