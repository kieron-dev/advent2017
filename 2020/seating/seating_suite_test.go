package seating_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSeating(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seating Suite")
}
