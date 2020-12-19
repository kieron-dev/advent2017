package rule_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rule Suite")
}
