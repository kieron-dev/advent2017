package two023_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2023(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2023 Suite")
}
