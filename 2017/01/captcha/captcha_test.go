package captcha_test

import (
	"github.com/kieron-pivotal/advent2017/01/captcha"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Captcha", func() {
	It("calcs captcha", func() {
		Expect(captcha.Decode("")).To(Equal(0))
		Expect(captcha.Decode("1212")).To(Equal(6))
		Expect(captcha.Decode("1221")).To(Equal(0))
		Expect(captcha.Decode("123425")).To(Equal(4))
		Expect(captcha.Decode("123123")).To(Equal(12))
		Expect(captcha.Decode("12131415")).To(Equal(4))
	})

})
