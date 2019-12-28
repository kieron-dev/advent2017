package manyworlds

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
	"github.com/kieron-pivotal/advent2017/advent2019/grid"
)

type World struct {
	grid     map[grid.Coord]rune
	keyCount int
	startPos grid.Coord
}

func NewWorld() *World {
	w := World{}
	w.grid = map[grid.Coord]rune{}

	return &w
}

func (w *World) LoadMap(r io.Reader) {
	fr := advent2019.FileReader{}

	row := 0
	fr.Each(r, func(line string) {
		for col, val := range line {
			coord := grid.NewCoord(col, row)
			w.grid[coord] = val
			if val == rune('@') {
				w.startPos = coord
				w.grid[coord] = rune('.')
			}
			if val >= 'a' && val <= 'z' {
				w.keyCount++
			}
		}
		row++
	})
}

func (w *World) KeysCount() int {
	return w.keyCount
}

func (w *World) StartPos() grid.Coord {
	return w.startPos
}

func (w *World) CharAt(c grid.Coord) rune {
	return w.grid[c]
}

type path struct {
	steps int
	keys  map[rune]bool
	pos   grid.Coord
}

func (p path) ToString() string {
	keySlice := []string{}
	for k := range p.keys {
		keySlice = append(keySlice, string(k))
	}
	sort.Strings(keySlice)
	keys := strings.Join(keySlice, "")
	return fmt.Sprintf("(%d,%d):%s", p.pos.X(), p.pos.Y(), keys)
}

func (w *World) MinStepsToCollectKeys() int {
	min := math.MaxInt32

	stack := []path{
		path{
			steps: 0,
			keys:  nil,
			pos:   w.startPos,
		},
	}

	visited := map[string]bool{}

	for len(stack) > 0 {
		curPath := stack[0]
		stack = stack[1:]

		curPathKey := curPath.ToString()
		if visited[curPathKey] {
			continue
		}
		visited[curPathKey] = true

		for _, dir := range []grid.Direction{grid.North, grid.South, grid.West, grid.East} {
			nextPos := curPath.pos.Move(dir)
			nextCell := w.grid[nextPos]
			switch {
			case nextCell == '.':
				stack = append(stack, path{
					steps: curPath.steps + 1,
					keys:  curPath.keys,
					pos:   nextPos,
				})
			case nextCell == '#':
				continue
			case nextCell >= 'a' && nextCell <= 'z':
				newKeys := map[rune]bool{}
				for k, v := range curPath.keys {
					newKeys[k] = v
				}
				newKeys[nextCell] = true
				if len(newKeys) == w.keyCount && curPath.steps+1 < min {
					min = curPath.steps + 1
					continue
				}
				stack = append(stack, path{
					steps: curPath.steps + 1,
					keys:  newKeys,
					pos:   nextPos,
				})
			case nextCell >= 'A' && nextCell <= 'Z':
				lower := 'a' + nextCell - 'A'
				if curPath.keys[lower] {
					stack = append(stack, path{
						steps: curPath.steps + 1,
						keys:  curPath.keys,
						pos:   nextPos,
					})
				}
			default:
				panic("didn't think we'd get here")
			}
		}
	}

	return min
}
