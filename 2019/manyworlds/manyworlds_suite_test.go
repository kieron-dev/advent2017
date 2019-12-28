package manyworlds_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestManyworlds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manyworlds Suite")
}
