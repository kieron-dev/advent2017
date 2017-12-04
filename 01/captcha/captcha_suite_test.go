package captcha_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCaptcha(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Captcha Suite")
}
