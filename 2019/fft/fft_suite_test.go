package fft_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFft(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fft Suite")
}
