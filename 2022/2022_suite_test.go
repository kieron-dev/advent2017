package two022_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2022(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2022 Suite")
}
