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

type graphNode struct {
	label  rune
	coord  grid.Coord
	linked map[rune]int
	dist   int
}

type World struct {
	grid       map[grid.Coord]rune
	keyCount   int
	startPos   grid.Coord
	graphNodes map[rune]*graphNode
}

func NewWorld() *World {
	w := World{}
	w.grid = map[grid.Coord]rune{}
	w.graphNodes = map[rune]*graphNode{}

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

	w.makeGraph()
}

func (w *World) getConnectedNodes(from *graphNode) []*graphNode {
	visited := map[grid.Coord]bool{}
	stack := []grid.Coord{from.coord}

	distances := map[grid.Coord]int{from.coord: 0}

	connected := []*graphNode{}

	for len(stack) > 0 {
		curPos := stack[0]
		stack = stack[1:]

		if visited[curPos] {
			continue
		}
		visited[curPos] = true

		if curPos != from.coord && w.grid[curPos] != '.' {
			connected = append(connected, &graphNode{
				coord: curPos,
				label: w.grid[curPos],
				dist:  distances[curPos],
			})
			continue
		}

		for _, dir := range []grid.Direction{grid.North, grid.South, grid.East, grid.West} {
			nextPos := curPos.Move(dir)
			cell := w.grid[nextPos]
			switch {
			case cell == '#':
				continue
			default:
				stack = append(stack, nextPos)
				d, ok := distances[nextPos]
				if !(ok && d > distances[curPos]+1) {
					distances[nextPos] = distances[curPos] + 1
				}
			}
		}
	}

	return connected
}

func (w *World) makeGraph() {
	start := &graphNode{
		label:  '0',
		coord:  w.startPos,
		linked: map[rune]int{},
	}

	stack := []*graphNode{start}

	visited := map[rune]bool{}

	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]

		if visited[cur.label] {
			continue
		}
		visited[cur.label] = true

		cur.linked = map[rune]int{}
		for _, conn := range w.getConnectedNodes(cur) {
			cur.linked[conn.label] = conn.dist
			stack = append(stack, conn)
		}

		w.graphNodes[cur.label] = cur
	}

}

func (w *World) PrintGraph() {
	for _, n := range w.graphNodes {
		fmt.Printf("node %q\n", n.label)
		for l, d := range n.linked {
			fmt.Printf("\t-> %q: %d\n", l, d)
		}
	}
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
	pos   rune
}

func (p path) ToString() string {
	keySlice := []string{}
	for k := range p.keys {
		keySlice = append(keySlice, string(k))
	}
	sort.Strings(keySlice)
	keys := strings.Join(keySlice, "")
	return fmt.Sprintf("%c:%s", p.pos, keys)
}

func (w *World) MinStepsToCollectKeys() int {
	min := math.MaxInt32

	stack := []path{
		path{
			steps: 0,
			keys:  nil,
			pos:   '0',
		},
	}

	visited := map[string]int{}

	for len(stack) > 0 {
		curPath := stack[0]
		stack = stack[1:]

		if curPath.steps >= min {
			continue
		}

		curPathKey := curPath.ToString()
		if n, ok := visited[curPathKey]; ok && n <= curPath.steps {
			continue
		}
		visited[curPathKey] = curPath.steps

		curNode := w.graphNodes[curPath.pos]
		keys := curPath.keys

		if curPath.pos >= 'a' && curPath.pos <= 'z' {
			keys = map[rune]bool{}
			for k, v := range curPath.keys {
				keys[k] = v
			}
			keys[curPath.pos] = true
		}

		if len(keys) == w.keyCount && curPath.steps < min {
			min = curPath.steps
		}

		for nextNode, dist := range curNode.linked {
			if nextNode >= 'A' && nextNode <= 'Z' {
				lower := 'a' + nextNode - 'A'
				if !keys[lower] {
					continue
				}
			}

			stack = append(stack, path{
				steps: curPath.steps + dist,
				keys:  keys,
				pos:   nextNode,
			})
		}
	}

	return min
}
