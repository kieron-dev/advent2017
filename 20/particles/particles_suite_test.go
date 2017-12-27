package particles_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParticles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Particles Suite")
}
