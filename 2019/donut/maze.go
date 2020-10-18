package donut

import (
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/kieron-dev/advent2017/advent2019"
	"github.com/kieron-dev/advent2017/advent2019/grid"
)

type Maze struct {
	layout         []string
	entrance, exit grid.Coord
	outerTeleports map[grid.Coord]grid.Coord
	innerTeleports map[grid.Coord]grid.Coord
}

func NewMaze() *Maze {
	m := Maze{}
	m.outerTeleports = map[grid.Coord]grid.Coord{}
	m.innerTeleports = map[grid.Coord]grid.Coord{}

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
	if to, ok := m.outerTeleports[from]; ok {
		return to, nil
	}
	if to, ok := m.innerTeleports[from]; ok {
		return to, nil
	}

	return grid.Coord{}, errors.New("not a teleport point")
}

type position struct {
	coord grid.Coord
	level int
}

func (m *Maze) ShortestPath(withLevels bool) int {
	stack := []position{position{coord: m.entrance, level: 0}}
	visited := map[position]bool{}
	distances := map[position]int{position{coord: m.entrance, level: 0}: 0}

	for len(stack) > 0 {
		curPoint := stack[0]
		stack = stack[1:]

		if visited[curPoint] {
			continue
		}
		visited[curPoint] = true

		if (!withLevels || curPoint.level == 0) && curPoint.coord == m.exit {
			return distances[curPoint]
		}

		for _, dir := range []grid.Direction{grid.North, grid.South, grid.West, grid.East} {
			nextCoord := curPoint.coord.Move(dir)
			cell := m.layout[nextCoord.Y()][nextCoord.X()]
			switch {
			case cell == '.':
				next := position{coord: nextCoord, level: curPoint.level}
				stack = append(stack, next)
				distances[next] = distances[curPoint] + 1
			case cell >= 'A' && cell <= 'Z':
				var next position
				if teleport, ok := m.outerTeleports[nextCoord]; ok && (!withLevels || curPoint.level != 0) {
					next = position{coord: teleport, level: curPoint.level - 1}
					stack = append(stack, next)
					distances[next] = distances[curPoint] + 1
				} else if teleport, ok := m.innerTeleports[nextCoord]; ok {
					next = position{coord: teleport, level: curPoint.level + 1}
					stack = append(stack, next)
					distances[next] = distances[curPoint] + 1
				}
			}
		}
	}
	return -1
}

type navType int

const (
	_ navType = iota
	inner
	outer
)

type gridPair struct {
	point    grid.Coord
	navPoint grid.Coord
	navType  navType
}

func newGridPair(px, py, npx, npy int, navType navType) gridPair {
	return gridPair{
		point:    grid.NewCoord(px, py),
		navPoint: grid.NewCoord(npx, npy),
		navType:  navType,
	}
}

func (m *Maze) analyse() {
	teleports := map[string][]gridPair{}

	re := regexp.MustCompile(`^ *([A-Z]{0,2})[.#]+([A-Z]{0,2})[ A-Z]*?([A-Z]{0,2})[.#]+([A-Z]{0,2}) *$`)

	for r, row := range m.layout {
		match := re.FindStringSubmatchIndex(row)
		if len(match) == 0 {
			continue
		}

		if match[3]-match[2] == 2 {
			key := fmt.Sprintf("%c%c", row[match[2]], row[match[2]+1])
			teleports[key] = append(teleports[key], newGridPair(match[2]+2, r, match[2]+1, r, outer))
		}
		if match[5]-match[4] == 2 {
			key := fmt.Sprintf("%c%c", row[match[4]], row[match[4]+1])
			teleports[key] = append(teleports[key], newGridPair(match[4]-1, r, match[4], r, inner))
		}
		if match[7]-match[6] == 2 {
			key := fmt.Sprintf("%c%c", row[match[6]], row[match[6]+1])
			teleports[key] = append(teleports[key], newGridPair(match[6]+2, r, match[6]+1, r, inner))
		}
		if match[9]-match[8] == 2 {
			key := fmt.Sprintf("%c%c", row[match[8]], row[match[8]+1])
			teleports[key] = append(teleports[key], newGridPair(match[8]-1, r, match[8], r, outer))
		}
	}

	for col := 0; col < len(m.layout[0]); col++ {
		line := ""
		for r := 0; r < len(m.layout); r++ {
			line += string(m.layout[r][col])
		}

		match := re.FindStringSubmatchIndex(line)
		if len(match) == 0 {
			continue
		}

		if match[3]-match[2] == 2 {
			key := fmt.Sprintf("%c%c", line[match[2]], line[match[2]+1])
			teleports[key] = append(teleports[key], newGridPair(col, match[2]+2, col, match[2]+1, outer))
		}
		if match[5]-match[4] == 2 {
			key := fmt.Sprintf("%c%c", line[match[4]], line[match[4]+1])
			teleports[key] = append(teleports[key], newGridPair(col, match[4]-1, col, match[4], inner))
		}
		if match[7]-match[6] == 2 {
			key := fmt.Sprintf("%c%c", line[match[6]], line[match[6]+1])
			teleports[key] = append(teleports[key], newGridPair(col, match[6]+2, col, match[6]+1, inner))
		}
		if match[9]-match[8] == 2 {
			key := fmt.Sprintf("%c%c", line[match[8]], line[match[8]+1])
			teleports[key] = append(teleports[key], newGridPair(col, match[8]-1, col, match[8], outer))
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
			if v[0].navType == outer {
				m.outerTeleports[v[0].navPoint] = v[1].point
			} else {
				m.innerTeleports[v[0].navPoint] = v[1].point
			}
			if v[1].navType == outer {
				m.outerTeleports[v[1].navPoint] = v[0].point
			} else {
				m.innerTeleports[v[1].navPoint] = v[0].point
			}
		}
	}
}
