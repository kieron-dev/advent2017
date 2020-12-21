// Package image finds monsters
package image

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	tile  []string
	edges []uint
	size  int
}

func NewTile(rows []string) *Tile {
	tile := Tile{
		tile: rows,
		size: len(rows[0]),
	}

	tile.setEdges()

	return &tile
}

func (t Tile) Content() []string {
	content := make([]string, t.size-2)
	for i := 1; i < t.size-1; i++ {
		content[i] = t.tile[i][1 : t.size-1]
	}

	return content
}

func (t Tile) LineContent(idx int) string {
	return t.tile[idx][1 : t.size-1]
}

func (t Tile) Print() {
	for _, s := range t.tile {
		fmt.Println(s)
	}
}

func (t *Tile) setEdges() {
	t.edges = []uint{
		EdgeToUint(t.tile[0]),
		EdgeToUint(t.RightEdge()),
		EdgeToUint(t.tile[t.size-1]),
		EdgeToUint(t.LeftEdge()),
	}
}

func (t Tile) Top() uint {
	return t.edges[0]
}

func (t Tile) Right() uint {
	return t.edges[1]
}

func (t Tile) Left() uint {
	return t.edges[3]
}

func (t Tile) Bottom() uint {
	return t.edges[2]
}

func (t *Tile) Rotate() {
	newRows := []string{}
	for i := 0; i < t.size; i++ {
		row := ""
		for j := t.size - 1; j >= 0; j-- {
			row += string(t.tile[j][i])
		}
		newRows = append(newRows, row)
	}

	t.tile = newRows
	t.setEdges()
}

func (t *Tile) flip() {
	newRows := []string{}

	for i := t.size - 1; i >= 0; i-- {
		newRows = append(newRows, t.tile[i])
	}

	t.tile = newRows
	t.setEdges()
}

var (
	body0RegExp = regexp.MustCompile(`..................#.`)
	body1RegExp = regexp.MustCompile(`#....##....##....###`)
	body2RegExp = regexp.MustCompile(`.#..#..#..#..#..#...`)
)

func (t Tile) MonsterCount() int {
	count := 0

	for i := 1; i < t.size-1; i++ {
		posLocs1 := body1RegExp.FindAllStringIndex(t.tile[i], -1)
		for _, positions := range posLocs1 {
			x := positions[0]
			if x+20 > t.size {
				continue
			}

			if body0RegExp.MatchString(t.tile[i-1][x:x+20]) &&
				body2RegExp.MatchString(t.tile[i+1][x:x+20]) {
				count++
			}
		}
	}

	return count
}

func (t *Tile) CountNonMonsterHashes() int {
	allHashes := 0

	for _, row := range t.tile {
		allHashes += strings.Count(row, "#")
	}

	return allHashes - 15*t.MonsterCount()
}

func (t *Tile) TransformTillMonster() {
	for i := 0; i < 4; i++ {
		if t.MonsterCount() > 0 {
			return
		}
		t.Rotate()
	}
	t.flip()
	for i := 0; i < 4; i++ {
		if t.MonsterCount() > 0 {
			return
		}
		t.Rotate()
	}

	log.Fatalf("oops - didn't find a monster")
}

func (t *Tile) TransformForLeft(leftEdge uint) {
	for i := 0; i < 4; i++ {
		if t.Left() == leftEdge {
			return
		}
		t.Rotate()
	}
	t.flip()
	for i := 0; i < 4; i++ {
		if t.Left() == leftEdge {
			return
		}
		t.Rotate()
	}

	log.Fatalf("oops")
}

func (t *Tile) TransformForTop(topEdge uint) {
	for i := 0; i < 4; i++ {
		if t.Top() == topEdge {
			return
		}
		t.Rotate()
	}
	t.flip()
	for i := 0; i < 4; i++ {
		if t.Top() == topEdge {
			return
		}
		t.Rotate()
	}

	log.Fatalf("oops")
}

func (t Tile) PossibleEdges() []uint {
	res := t.edges[:]
	for _, edge := range t.edges {
		res = append(res, t.ReverseEdge(edge))
	}

	return res
}

func EdgeToUint(edge string) uint {
	var u uint

	for _, c := range edge {
		u <<= 1
		if c == rune('#') {
			u |= uint(1)
		}
	}

	return u
}

func (t Tile) ReverseEdge(e uint) uint {
	var u uint

	for i := 0; i < t.size; i++ {
		u <<= 1
		u |= (e & 1)
		e >>= 1
	}

	return u
}

func (t Tile) LeftEdge() string {
	var s string

	for _, line := range t.tile {
		s += string(line[0])
	}

	return s
}

func (t Tile) RightEdge() string {
	var s string

	for _, line := range t.tile {
		s += string(line[len(line)-1])
	}

	return s
}

type Sorter struct {
	tiles         map[int]*Tile
	edgeLocations map[uint][]int
}

func NewSorter() Sorter {
	return Sorter{
		tiles:         map[int]*Tile{},
		edgeLocations: map[uint][]int{},
	}
}

var tileRE = regexp.MustCompile(`Tile (\d+):`)

func (s *Sorter) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	var rows []string
	tileNo := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := tileRE.FindStringSubmatch(line)
		if len(matches) == 2 {
			if rows != nil {
				s.tiles[tileNo] = NewTile(rows)
				rows = nil
			}

			var err error
			tileNo, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("getting tile no failed: %v", err)
			}

			continue
		}

		rows = append(rows, line)
	}

	s.tiles[tileNo] = NewTile(rows)

	s.registerEdges()
}

func (s *Sorter) registerEdges() {
	for tileNo, tile := range s.tiles {
		for _, edge := range tile.PossibleEdges() {
			s.edgeLocations[edge] = append(s.edgeLocations[edge], tileNo)
		}
	}
}

func (s Sorter) CornerProduct() int {
	prod := 1
	for _, tileNo := range s.Corners() {
		prod *= tileNo
	}

	return prod
}

func (s *Sorter) Solve() int {
	bigTile := s.mergeImages()

	bigTile.TransformTillMonster()

	return bigTile.CountNonMonsterHashes()
}

func (s *Sorter) mergeImages() *Tile {
	rows := [][]int{}

	corners := s.Corners()
	firstCorner := corners[0]
	topLeft := s.tiles[firstCorner]

	for len(s.edgeLocations[topLeft.Top()]) > 1 || len(s.edgeLocations[topLeft.Left()]) > 1 {
		topLeft.Rotate()
	}

	cur := topLeft
	taken := map[int]bool{firstCorner: true}
	row := []int{firstCorner}

	for len(s.edgeLocations[cur.Right()]) > 1 {
		possibilities := []int{}
		for _, tnum := range s.edgeLocations[cur.Right()] {
			if !taken[tnum] {
				possibilities = append(possibilities, tnum)
			}
		}

		if len(possibilities) > 1 {
			log.Fatalf("need more work then!")
		}

		needLeft := cur.Right()

		cur = s.tiles[possibilities[0]]
		cur.TransformForLeft(needLeft)
		taken[possibilities[0]] = true
		row = append(row, possibilities[0])
	}

	rows = append(rows, row)

	lastRow := 0
	for len(taken) < len(s.tiles) {
		row = []int{}

		for i := 0; i < len(rows[0]); i++ {
			possibilities := []int{}
			above := s.tiles[rows[lastRow][i]]
			for _, tnum := range s.edgeLocations[above.Bottom()] {
				if !taken[tnum] {
					possibilities = append(possibilities, tnum)
				}
			}

			if len(possibilities) > 1 {
				log.Fatalf("need more work then!")
			}

			needTop := above.Bottom()

			found := s.tiles[possibilities[0]]
			found.TransformForTop(needTop)
			taken[possibilities[0]] = true
			row = append(row, possibilities[0])

		}

		rows = append(rows, row)
		lastRow++
	}

	bigTile := []string{}

	for _, row := range rows {
		for i := 1; i < topLeft.size-1; i++ {
			line := ""
			for _, tileNo := range row {
				line += s.tiles[tileNo].LineContent(i)
			}
			bigTile = append(bigTile, line)
		}
	}

	return NewTile(bigTile)
}

func (s Sorter) Corners() []int {
	tileUniqueCounts := map[int]int{}

	for edge, locations := range s.edgeLocations {
		tile := s.tiles[locations[0]]
		if len(locations) == 1 && len(s.edgeLocations[tile.ReverseEdge(edge)]) == 1 {
			tileUniqueCounts[locations[0]]++
		}
	}

	corners := []int{}
	for tileNo, count := range tileUniqueCounts {
		if count == 4 {
			corners = append(corners, tileNo)
		}
	}

	return corners
}
