package jolt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJolt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jolt Suite")
}
