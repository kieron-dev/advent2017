package tiled_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/tiled"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Floor", func() {
	var (
		data  io.Reader
		floor tiled.Floor
	)

	BeforeEach(func() {
		data = strings.NewReader(`
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
`)
		floor = tiled.NewFloor()
		floor.Load(data)
	})

	It("can do the example", func() {
		Expect(floor.BlackCount()).To(Equal(10))
	})

	It("can evolve", func() {
		floor.Evolve()
		Expect(floor.BlackCount()).To(Equal(15))
		floor.Evolve()
		Expect(floor.BlackCount()).To(Equal(12))
		floor.Evolve()
		Expect(floor.BlackCount()).To(Equal(25))
	})

	It("can evolve a lot", func() {
		for i := 0; i < 100; i++ {
			floor.Evolve()
		}
		Expect(floor.BlackCount()).To(Equal(2208))
	})
})
