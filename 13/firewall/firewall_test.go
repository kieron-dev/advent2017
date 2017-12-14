package firewall_test

import (
	"github.com/kieron-pivotal/advent2017/13/firewall"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Firewall", func() {

	Describe("Detector", func() {
		It("knows position at time t", func() {
			Expect(firewall.DetectorPosition(2, 0)).To(Equal(0))
			Expect(firewall.DetectorPosition(2, 1)).To(Equal(1))
			Expect(firewall.DetectorPosition(2, 2)).To(Equal(0))
			Expect(firewall.DetectorPosition(2, 3)).To(Equal(1))

			Expect(firewall.DetectorPosition(3, 0)).To(Equal(0))
			Expect(firewall.DetectorPosition(3, 1)).To(Equal(1))
			Expect(firewall.DetectorPosition(3, 2)).To(Equal(2))
			Expect(firewall.DetectorPosition(3, 3)).To(Equal(1))
			Expect(firewall.DetectorPosition(3, 4)).To(Equal(0))
			Expect(firewall.DetectorPosition(3, 10)).To(Equal(2))

			Expect(firewall.DetectorPosition(4, 0)).To(Equal(0))
			Expect(firewall.DetectorPosition(4, 1)).To(Equal(1))
			Expect(firewall.DetectorPosition(4, 2)).To(Equal(2))
			Expect(firewall.DetectorPosition(4, 3)).To(Equal(3))
			Expect(firewall.DetectorPosition(4, 4)).To(Equal(2))
			Expect(firewall.DetectorPosition(4, 5)).To(Equal(1))
			Expect(firewall.DetectorPosition(4, 6)).To(Equal(0))
		})
	})

	It("calcs severity", func() {
		Expect(firewall.Severity(map[int]int{0: 3, 1: 2, 4: 4, 6: 4}, 0)).To(Equal(24))
	})

	It("calcs min delay", func() {
		Expect(firewall.MinDelay(map[int]int{0: 3, 1: 2, 4: 4, 6: 4})).To(Equal(10))
	})

})
