package two023_test

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type instruction struct {
	dir  direction
	dist int
}

type lavaduct struct {
	cells            map[coord]bool
	rows, cols       map[int]bool
	rowList, colList []int
}

var dirsA = map[string]direction{
	"U": up,
	"D": down,
	"L": left,
	"R": right,
}

var dirsB = map[string]direction{
	"3": up,
	"1": down,
	"2": left,
	"0": right,
}

func translateInstructionsA(r io.Reader) []instruction {
	ret := []instruction{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		bits := strings.Split(line, " ")
		dir := dirsA[bits[0]]
		dist, err := strconv.Atoi(bits[1])
		Expect(err).NotTo(HaveOccurred())
		ret = append(ret, instruction{dir: dir, dist: dist})
	}
	return ret
}

func translateInstructionsB(r io.Reader) []instruction {
	ret := []instruction{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		b1 := strings.Index(line, "(")
		b2 := strings.Index(line, ")")

		dist, err := strconv.ParseInt(line[b1+2:b2-1], 16, 32)
		Expect(err).NotTo(HaveOccurred())
		dir := dirsB[line[b2-1:b2]]
		ret = append(ret, instruction{dir: dir, dist: int(dist)})
	}
	return ret
}

func calcArea(instructions []instruction) int {
	l := lavaduct{}
	l.rows = map[int]bool{}
	l.cols = map[int]bool{}

	start := coord{0, 0}
	l.cells = map[coord]bool{start: true}

	pos := start
	for _, instr := range instructions {
		pos = pos.add(coord(instr.dir).mult(instr.dist))
		l.rows[pos[0]] = true
		l.cols[pos[1]] = true
	}

	l.recordPoints()

	condensedLavaduct, rowMap, colMap := l.newCondensed(instructions)
	revRowMap := map[int]int{}
	var maxCol, maxRow int
	for k, v := range rowMap {
		revRowMap[v] = k
		if v > maxRow {
			maxRow = v
		}
	}
	revColMap := map[int]int{}
	for k, v := range colMap {
		revColMap[v] = k
		if v > maxCol {
			maxCol = v
		}
	}

	sum := 0
	for r := 0; r <= maxRow; r += 2 {
		for c := 0; c <= maxCol; c += 2 {
			if condensedLavaduct.cells[coord{r + 1, c + 1}] {
				sum += (revRowMap[r+2] - revRowMap[r]) * (revColMap[c+2] - revColMap[c])
				if !condensedLavaduct.cells[coord{r + 1, c + 3}] {
					sum += (revRowMap[r+2] - revRowMap[r])
				}
				if !condensedLavaduct.cells[coord{r + 3, c + 1}] {
					sum += (revColMap[c+2] - revColMap[c])
				}
			}
		}
	}

	return sum + 1
}

func (l lavaduct) newCondensed(instructions []instruction) (lavaduct, map[int]int, map[int]int) {
	newLavaduct := lavaduct{}
	newLavaduct.rows = map[int]bool{}
	newLavaduct.cols = map[int]bool{}
	newLavaduct.cells = map[coord]bool{}

	rowMap := map[int]int{}
	colMap := map[int]int{}

	for i, r := range l.rowList {
		rowMap[r] = i * 2
	}
	for i, c := range l.colList {
		colMap[c] = i * 2
	}

	posOrig := coord{0, 0}
	posMapped := coord{rowMap[0], colMap[0]}

	for _, instr := range instructions {

		var startOrig, startMapped, endOrig, endMapped int

		if instr.dir == up || instr.dir == down {
			startMapped = posMapped[0]
			startOrig = posOrig[0]

			// down
			endOrig = startOrig + instr.dist
			if instr.dir == up {
				endOrig = startOrig - instr.dist
			}
			endMapped = rowMap[endOrig]
		}

		if instr.dir == left || instr.dir == right {
			startMapped = posMapped[1]
			startOrig = posOrig[1]

			// right
			endOrig := startOrig + instr.dist
			if instr.dir == left {
				endOrig = startOrig - instr.dist
			}
			endMapped = colMap[endOrig]
		}

		for i := 1; i <= absDiff(startMapped, endMapped); i++ {
			posMapped = posMapped.add(coord(instr.dir))
			newLavaduct.cells[posMapped] = true
			newLavaduct.rows[posMapped[0]] = true
			newLavaduct.cols[posMapped[1]] = true
		}

		posOrig = posOrig.add(coord(instr.dir).mult(instr.dist))
	}

	newLavaduct.recordPoints()
	newLavaduct.fill()

	return newLavaduct, rowMap, colMap
}

func (l *lavaduct) recordPoints() {
	for k := range l.rows {
		l.rowList = append(l.rowList, k)
	}
	slices.Sort(l.rowList)
	for k := range l.cols {
		l.colList = append(l.colList, k)
	}
	slices.Sort(l.colList)
}

// func (l lavaduct) print() {
// 	rowMin := slices.Min(l.rowList)
// 	rowMax := slices.Max(l.rowList)
// 	colMin := slices.Min(l.colList)
// 	colMax := slices.Max(l.colList)
// 	for r := rowMin; r <= rowMax; r++ {
// 		for c := colMin; c <= colMax; c++ {
// 			if l.cells[coord{r, c}] {
// 				fmt.Printf("#")
// 			} else {
// 				fmt.Printf(".")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func (l lavaduct) interiorPoint() coord {
	rowMin := slices.Min(l.rowList)
	rowMax := slices.Max(l.rowList)
	colMin := slices.Min(l.colList)
	colMax := slices.Max(l.colList)
outer:
	for r := rowMin; r <= rowMax; r++ {
		for c := colMin; c < colMax; c++ {
			if l.cells[coord{r, c}] {
				if !l.cells[coord{r, c + 1}] {
					return coord{r, c + 1}
				}
				continue outer
			}
		}
	}
	Fail("didn't find an interior point")
	return coord{}
}

func (l lavaduct) fill() {
	from := l.interiorPoint()
	added := map[coord]bool{}
	queue := []coord{from}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if l.cells[cur] {
			continue
		}
		l.cells[cur] = true

		for _, d := range []direction{up, down, left, right} {
			next := cur.add(coord(d))
			if added[next] {
				continue
			}
			added[next] = true
			queue = append(queue, next)
		}
	}
}

var _ = Describe("18", func() {
	It("does part A", func() {
		f, err := os.Open("input18")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		instructions := translateInstructionsA(f)

		l := calcArea(instructions)
		Expect(l).To(Equal(48795))
	})

	It("does part B", func() {
		f, err := os.Open("input18")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		instructions := translateInstructionsB(f)

		l := calcArea(instructions)
		Expect(l).To(Equal(40654918441248))
	})
})
