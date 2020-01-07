package springdroid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpringdroid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Springdroid Suite")
}
