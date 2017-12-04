package captcha_test

import (
	"github.com/kieron-pivotal/advent2017/01/captcha"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Captcha", func() {
	It("calcs captcha", func() {
		Expect(captcha.Decode("")).To(Equal(0))
		Expect(captcha.Decode("1122")).To(Equal(3))
		Expect(captcha.Decode("1111")).To(Equal(4))
	})

})
