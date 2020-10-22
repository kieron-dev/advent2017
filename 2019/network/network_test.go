package network_test

import (
	"io/ioutil"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-dev/advent2017/advent2019/network"
)

var _ = Describe("Network", func() {
	var input []byte

	BeforeEach(func() {
		var err error
		input, err = ioutil.ReadFile("../days/input23")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("seeing what a single NIC does", func() {
		var nic *network.NIC

		BeforeEach(func() {
			nic = network.NewNIC()

			nic.Start(strings.TrimSpace(string(input)))
			nic.Input(3)
			nic.Input(-1)
		})

		It("does something", func() {
			Eventually(nic.Output()).Should(Receive())
		})
	})

	Context("Day 23 part A", func() {
		var net *network.TokenRing

		BeforeEach(func() {
			net = network.NewNetwork(50, strings.TrimSpace(string(input)))
		})

		It("does something", func() {
			net.Start()
			special := net.SpecialQueueHead()
			Expect(special).To(HaveLen(2))
			Expect(special[1]).To(Equal(14834))
		})
	})

	Context("Day 23 part B", func() {
		var net *network.TokenRing

		BeforeEach(func() {
			net = network.NewNetwork(50, strings.TrimSpace(string(input)))
		})

		It("does something", func() {
			yVals := map[int]bool{}
			var out int

			for {
				net.Start()
				packet := net.SpecialQueueTail()
				if yVals[packet[1]] {
					out = packet[1]
					break
				}
				yVals[packet[1]] = true

				net.Enqueue(255, 0, packet[0], packet[1])
			}

			Expect(out).To(Equal(10215))
		})
	})
})
