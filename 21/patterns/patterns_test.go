package patterns_test

import (
	"strings"

	"github.com/kieron-pivotal/advent2017/21/patterns"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Patterns", func() {

	var (
		sq3x3 [][]byte
		sq2x2 [][]byte
	)

	BeforeEach(func() {
		sq3x3 = [][]byte{
			{'.', '#', '.'},
			{'.', '.', '#'},
			{'#', '#', '#'},
		}
		sq2x2 = [][]byte{
			{'1', '2'},
			{'3', '4'},
		}
	})

	Context("patterns", func() {
		It("creates a new pattern with the standard format", func() {
			art := patterns.New()
			Expect(art.Size()).To(Equal(3))
			Expect(art.Pattern()).To(Equal([]string{
				".#.",
				"..#",
				"###",
			}))
		})

		It("gets the 0,0 square", func() {
			art := patterns.New()
			Expect(art.GetSquare(0, 0)).To(Equal(sq3x3))
		})

		It("can rotate squares once clockwise", func() {
			patterns.RotateSquare(sq3x3)
			Expect(sq3x3).To(Equal([][]byte{
				{'#', '.', '.'},
				{'#', '.', '#'},
				{'#', '#', '.'},
			}))

			patterns.RotateSquare(sq2x2)
			Expect(sq2x2).To(Equal([][]byte{
				{'3', '1'},
				{'4', '2'},
			}))
		})

		It("can flip a square vertically", func() {
			patterns.FlipSquare(sq3x3)
			Expect(sq3x3).To(Equal([][]byte{
				{'.', '#', '.'},
				{'#', '.', '.'},
				{'#', '#', '#'},
			}))

			patterns.FlipSquare(sq2x2)
			Expect(sq2x2).To(Equal([][]byte{
				{'2', '1'},
				{'4', '3'},
			}))
		})

		It("can convert square to text", func() {
			Expect(patterns.SquareToString(sq3x3)).To(Equal(".#./..#/###"))
			Expect(patterns.SquareToString(sq2x2)).To(Equal("12/34"))
		})

		It("gets 3x3 patterns at coords", func() {
			art := patterns.New()
			keys := art.GetKeys(0, 0)
			Expect(keys).To(Equal([]string{
				".#./..#/###",
				"#../#.#/##.",
				"###/#../.#.",
				".##/#.#/..#",
				".#./#../###",
				"##./#.#/#..",
				"###/..#/.#.",
				"..#/#.#/.##",
			}))
		})
	})

	Context("rules", func() {
		It("can find matching rule", func() {
			art := patterns.New()
			newPattern := "abcd/abcd/abcd/abcd"
			art.AddRule("#../#.#/##.", newPattern)
			Expect(art.GetNewPattern(0, 0)).To(Equal(strings.Split(newPattern, "/")))
		})

		It("can progress one iteration", func() {
			art := patterns.New()
			newPattern := "abcd/abcd/abcd/abcd"
			art.AddRule("#../#.#/##.", newPattern)
			art.Advance()
			Expect(art.Pattern()).To(Equal([]string{
				"abcd",
				"abcd",
				"abcd",
				"abcd",
			}))
		})

		It("can progress two iterations", func() {
			art := patterns.New()
			art.AddRule("../.#", "##./#../...")
			art.AddRule(".#./..#/###", "#..#/..../..../#..#")
			art.Advance()
			art.Advance()
			Expect(art.Pattern()).To(Equal([]string{
				"##.##.",
				"#..#..",
				"......",
				"##.##.",
				"#..#..",
				"......",
			}))
		})

		It("can count on cells", func() {
			art := patterns.New()
			art.AddRule("../.#", "##./#../...")
			art.AddRule(".#./..#/###", "#..#/..../..../#..#")
			art.Advance()
			art.Advance()
			Expect(art.OnCount()).To(Equal(12))
		})
	})

})
