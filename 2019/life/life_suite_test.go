package life_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Life Suite")
}
