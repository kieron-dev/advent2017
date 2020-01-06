package donut_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/donut"
	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Donut", func() {

	var (
		d      *donut.Maze
		layout io.Reader
	)

	BeforeEach(func() {
		d = donut.NewMaze()
	})

	JustBeforeEach(func() {
		d.Load(layout)
	})

	Context("basic", func() {
		BeforeEach(func() {
			layout = strings.NewReader(`         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       
`)
		})

		It("can find the entrance", func() {
			Expect(d.Entrance()).To(Equal(grid.NewCoord(9, 2)))
		})

		It("can find the exit", func() {
			Expect(d.Exit()).To(Equal(grid.NewCoord(13, 16)))
		})

		It("links BC teleports", func() {
			Expect(d.Teleport(grid.NewCoord(1, 8))).To(Equal(grid.NewCoord(9, 6)))
			Expect(d.Teleport(grid.NewCoord(9, 7))).To(Equal(grid.NewCoord(2, 8)))
		})

		It("recognises non-teleports", func() {
			_, err := d.Teleport(grid.NewCoord(1, 1))
			Expect(err).To(MatchError("not a teleport point"))
		})

		It("has a shortest path of 23", func() {
			Expect(d.ShortestPath(false)).To(Equal(23))
		})

		It("has a recursive shortest path of 26", func() {
			Expect(d.ShortestPath(true)).To(Equal(26))
		})
	})

	Context("harder", func() {
		BeforeEach(func() {
			layout = strings.NewReader(`                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `)
		})

		It("has a shortest path of 58", func() {
			Expect(d.ShortestPath(false)).To(Equal(58))
		})
	})

	Context("hard recursive", func() {
		BeforeEach(func() {
			layout = strings.NewReader(`             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     
`)
		})

		It("has a shortest path of 396", func() {
			Expect(d.ShortestPath(true)).To(Equal(396))
		})
	})
})
