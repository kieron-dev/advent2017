package passphrase_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPassphrase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passphrase Suite")
}
