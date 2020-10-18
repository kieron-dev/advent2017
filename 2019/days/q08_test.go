package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q08", func() {
	var (
		bytes []byte
		image *advent2019.Image
	)

	BeforeEach(func() {
		var err error
		bytes, err = ioutil.ReadFile("./input08")
		if err != nil {
			panic(err)
		}
		image = advent2019.NewImage(6, 25)
	})

	It("does part A", func() {
		image.Load(strings.TrimSpace(string(bytes)))

		layer := image.FindLayerWithFewestZeros()
		res := layer.Count(1) * layer.Count(2)

		Expect(res).To(Equal(1848))
	})

	It("does part B", func() {
		image.Load(strings.TrimSpace(string(bytes)))

		layer := image.Decode()
		Expect(layer.String()).To(Equal(`XXXX  XX    XX X  X XXXX 
X    X  X    X X  X    X 
XXX  X       X X  X   X  
X    X XX    X X  X  X   
X    X  X X  X X  X X    
X     XXX  XX   XX  XXXX 
`))
	})
})
