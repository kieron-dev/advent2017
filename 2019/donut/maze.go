package donut

import (
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/kieron-pivotal/advent2017/advent2019"
	"github.com/kieron-pivotal/advent2017/advent2019/grid"
)

type Maze struct {
	layout         []string
	entrance, exit grid.Coord
	teleports      map[grid.Coord]grid.Coord
}

func NewMaze() *Maze {
	m := Maze{}
	m.teleports = map[grid.Coord]grid.Coord{}

	return &m
}

func (m *Maze) Load(layout io.Reader) {
	fr := advent2019.FileReader{}

	fr.Each(layout, func(line string) {
		m.layout = append(m.layout, line)
	})

	m.analyse()
}

func (m *Maze) Entrance() grid.Coord {
	return m.entrance
}

func (m *Maze) Exit() grid.Coord {
	return m.exit
}

func (m *Maze) Teleport(from grid.Coord) (grid.Coord, error) {
	to, ok := m.teleports[from]
	if !ok {
		return grid.Coord{}, errors.New("not a teleport point")
	}
	return to, nil
}

func (m *Maze) ShortestPath() int {
	stack := []grid.Coord{m.entrance}
	visited := map[grid.Coord]bool{}
	distances := map[grid.Coord]int{m.entrance: 0}

	for len(stack) > 0 {
		curPoint := stack[0]
		stack = stack[1:]

		if visited[curPoint] {
			continue
		}
		visited[curPoint] = true

		if curPoint == m.exit {
			return distances[curPoint]
		}

		for _, dir := range []grid.Direction{grid.North, grid.South, grid.West, grid.East} {
			next := curPoint.Move(dir)
			cell := m.layout[next.Y()][next.X()]
			switch {
			case cell == '.':
				stack = append(stack, next)
				distances[next] = distances[curPoint] + 1
			case cell >= 'A' && cell <= 'Z':
				if teleport, ok := m.teleports[next]; ok {
					stack = append(stack, teleport)
					distances[teleport] = distances[curPoint] + 1
				}
			}
		}
	}
	return -1
}

type gridPair struct {
	point    grid.Coord
	navPoint grid.Coord
}

func newGridPair(px, py, npx, npy int) gridPair {
	return gridPair{
		point:    grid.NewCoord(px, py),
		navPoint: grid.NewCoord(npx, npy),
	}
}

func (m *Maze) analyse() {
	teleports := map[string][]gridPair{}

	leftRE := regexp.MustCompile(`([A-Z][A-Z])\.`)
	rightRE := regexp.MustCompile(`\.([A-Z][A-Z])`)

	for r, row := range m.layout {
		matches := leftRE.FindAllStringSubmatchIndex(row, -1)
		for _, match := range matches {
			key := fmt.Sprintf("%c%c", row[match[0]], row[match[0]+1])
			teleports[key] = append(teleports[key], newGridPair(match[0]+2, r, match[0]+1, r))
		}

		matches = rightRE.FindAllStringSubmatchIndex(row, -1)
		for _, match := range matches {
			key := fmt.Sprintf("%c%c", row[match[0]+1], row[match[0]+2])
			teleports[key] = append(teleports[key], newGridPair(match[0], r, match[0]+1, r))
		}
	}

	for col := 0; col < len(m.layout[0]); col++ {
		line := ""
		for r := 0; r < len(m.layout); r++ {
			line += string(m.layout[r][col])
		}

		matches := leftRE.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			key := fmt.Sprintf("%c%c", line[match[0]], line[match[0]+1])
			teleports[key] = append(teleports[key], newGridPair(col, match[0]+2, col, match[0]+1))
		}

		matches = rightRE.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			key := fmt.Sprintf("%c%c", line[match[0]+1], line[match[0]+2])
			teleports[key] = append(teleports[key], newGridPair(col, match[0], col, match[0]+1))
		}
	}

	for k, v := range teleports {
		switch k {
		case "AA":
			if len(v) != 1 {
				panic("entrance")
			}
			m.entrance = v[0].point
		case "ZZ":
			if len(v) != 1 {
				panic("exit")
			}
			m.exit = v[0].point
		default:
			if len(v) != 2 {
				panic("teleport: " + k)
			}
			m.teleports[v[0].navPoint] = v[1].point
			m.teleports[v[1].navPoint] = v[0].point
		}
	}
}
