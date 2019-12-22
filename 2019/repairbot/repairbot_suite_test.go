package repairbot_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRepairbot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repairbot Suite")
}
