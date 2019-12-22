package vacuumbot_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVacuumbot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vacuumbot Suite")
}
