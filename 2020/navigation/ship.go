package navigation

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/maps"
)

var (
	north   = maps.NewVector(0, -1)
	south   = maps.NewVector(0, 1)
	west    = maps.NewVector(-1, 0)
	east    = maps.NewVector(1, 0)
	compass = []maps.Vector{north, east, south, west}
)

type Action int

const (
	Unknown Action = iota
	GoNorth
	GoSouth
	GoWest
	GoEast
	GoForward
	RotateLeft
	RotateRight
)

var actionMapping = map[string]Action{
	"N": GoNorth,
	"S": GoSouth,
	"E": GoEast,
	"W": GoWest,
	"F": GoForward,
	"L": RotateLeft,
	"R": RotateRight,
}

type Instruction struct {
	action Action
	amount int
}

func parseInstruction(inst string) Instruction {
	action := actionMapping[inst[:1]]
	if action == Unknown {
		log.Fatalf("unknown action for %q", inst)
	}

	n, err := strconv.Atoi(inst[1:])
	if err != nil {
		log.Fatalf("unknown num for %q", inst)
	}

	return Instruction{
		action: action,
		amount: n,
	}
}

type Ship struct {
	position       maps.Coord
	directionIndex int
	instructions   []Instruction
	waypoint       maps.Vector
}

func NewShip() Ship {
	return Ship{
		directionIndex: 1,
		waypoint:       maps.NewVector(10, -1),
	}
}

func (s *Ship) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		s.instructions = append(s.instructions, parseInstruction(line))
	}
}

func (s *Ship) Move() {
	for _, i := range s.instructions {
		switch i.action {
		case GoNorth:
			s.position = s.position.Plus(north.Times(i.amount))
		case GoSouth:
			s.position = s.position.Plus(south.Times(i.amount))
		case GoWest:
			s.position = s.position.Plus(west.Times(i.amount))
		case GoEast:
			s.position = s.position.Plus(east.Times(i.amount))
		case GoForward:
			s.position = s.position.Plus(compass[s.directionIndex].Times(i.amount))
		case RotateLeft:
			s.directionIndex += 4 - i.amount/90
			s.directionIndex = s.directionIndex % 4
		case RotateRight:
			s.directionIndex += i.amount / 90
			s.directionIndex = s.directionIndex % 4
		default:
			log.Fatalf("wtf happened? - %v", i)
		}
	}
}

func (s *Ship) MoveNew() {
	for _, i := range s.instructions {
		switch i.action {
		case GoNorth:
			s.waypoint = s.waypoint.Plus(north.Times(i.amount))
		case GoSouth:
			s.waypoint = s.waypoint.Plus(south.Times(i.amount))
		case GoWest:
			s.waypoint = s.waypoint.Plus(west.Times(i.amount))
		case GoEast:
			s.waypoint = s.waypoint.Plus(east.Times(i.amount))
		case GoForward:
			s.position = s.position.Plus(s.waypoint.Times(i.amount))
		case RotateLeft:
			for j := 0; j < i.amount; j += 90 {
				s.waypoint = s.waypoint.RotateLeft()
			}
		case RotateRight:
			for j := 0; j < i.amount; j += 90 {
				s.waypoint = s.waypoint.RotateRight()
			}
		default:
			log.Fatalf("wtf happened? - %v", i)
		}
	}
}

func (s Ship) ManhattanDistance() int {
	return mod(s.position.X) + mod(s.position.Y)
}

func mod(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
