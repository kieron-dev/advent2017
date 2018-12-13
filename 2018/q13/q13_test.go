package q13_test

import (
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q13"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q13", func() {

	var (
		f *os.File
	)

	BeforeEach(func() {
		var err error
		f, err = os.Open("example")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("loading a model", func() {
		It("can load a map", func() {
			mine := q13.NewMine(f)
			Expect(mine.Map).To(HaveLen(6))
			Expect(mine.Map[1]).To(Equal(`|   |  /----\`))
		})

		It("replaces carts with correct track", func() {
			mine := q13.NewMine(f)
			Expect(mine.Map[0]).To(Equal(`/---\        `))
			Expect(mine.Map[3]).To(Equal(`| | |  | |  |`))
		})

		It("records carts and directions", func() {
			mine := q13.NewMine(f)
			Expect(mine.Carts).To(HaveLen(2))
			Expect(*mine.Carts[0]).To(Equal(q13.Cart{Row: 0, Col: 2, Dir: q13.Right}))
			Expect(*mine.Carts[1]).To(Equal(q13.Cart{Row: 3, Col: 9, Dir: q13.Down}))
		})
	})

	Context("moving", func() {
		It("can sort carts", func() {
			carts := []*q13.Cart{
				{Row: 2, Col: 5},
				{Row: 1, Col: 5},
				{Row: 3, Col: 2},
				{Row: 3, Col: 1},
			}
			m := q13.Mine{Carts: carts}
			m.SortCarts()
			Expect(m.Carts[0].Row).To(Equal(1))
			Expect(m.Carts[1].Row).To(Equal(2))
			Expect(m.Carts[2].Col).To(Equal(1))
			Expect(m.Carts[3].Col).To(Equal(2))
		})
	})

	It("can move cars", func() {
		mine := q13.NewMine(f)
		mine.MoveCarts()
		Expect(mine.Carts[0].Row).To(Equal(0))
		Expect(mine.Carts[0].Col).To(Equal(3))
		Expect(mine.Carts[0].Dir).To(Equal(q13.Right))

		Expect(mine.Carts[1].Row).To(Equal(4))
		Expect(mine.Carts[1].Col).To(Equal(9))
		Expect(mine.Carts[1].Dir).To(Equal(q13.Right))
	})

	It("can detect a crash", func() {
		mine := q13.NewMine(f)
		i, crashedCart := mine.RunTillCrash()
		Expect(i).To(Equal(14))
		Expect(crashedCart.Row).To(Equal(3))
		Expect(crashedCart.Col).To(Equal(7))
	})
})
