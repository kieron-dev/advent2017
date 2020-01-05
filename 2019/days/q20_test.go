package days_test

import (
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/donut"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q20", func() {
	var (
		d *donut.Maze
	)

	BeforeEach(func() {
		d = donut.NewMaze()
		f, err := os.Open("input20")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		d.Load(f)
	})

	FIt("does part A", func() {
		Expect(d.ShortestPath()).To(Equal(644))
	})

})
