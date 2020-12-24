package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/cards"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("22", func() {
	var (
		data   *os.File
		combat cards.Combat
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input22")
		Expect(err).NotTo(HaveOccurred())

		combat = cards.NewCombat()
		combat.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		for combat.Winner() == 0 {
			combat.Play()
		}

		Expect(combat.Score(combat.Winner())).To(Equal(32824))
	})

	It("does part B", func() {
		combat.SetRecursive(true)

		for combat.Winner() == 0 {
			combat.Play()
		}

		Expect(combat.Score(combat.Winner())).To(Equal(36515))
	})
})
