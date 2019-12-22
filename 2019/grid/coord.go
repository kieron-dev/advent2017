package grid

type Direction int

const (
	_ Direction = iota
	North
	South
	West
	East
)

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("not a direction")
	}
}

type Coord struct {
	x int
	y int
}

func NewCoord(x, y int) Coord {
	return Coord{x: x, y: y}
}

func (c Coord) X() int {
	return c.x
}

func (c Coord) Y() int {
	return c.y
}

func (c Coord) Minus(d Coord) Coord {
	return Coord{
		x: c.x - d.x,
		y: c.y - d.y,
	}
}

func (c Coord) Mag2() int {
	return c.x*c.x + c.y*c.y
}

func (c Coord) quadrant() int {
	if c.x >= 0 && c.y <= 0 {
		return 0
	}
	if c.x >= 0 && c.y >= 0 {
		return 1
	}
	if c.x <= 0 && c.y >= 0 {
		return 2
	}
	if c.x <= 0 && c.y <= 0 {
		return 3
	}
	panic(c)
}

func (c Coord) Add(d Coord) Coord {
	return Coord{
		x: c.x + d.x,
		y: c.y + d.y,
	}
}

func (c Coord) Move(d Direction) Coord {
	switch d {
	case North:
		return Coord{
			x: c.x,
			y: c.y - 1,
		}
	case South:
		return Coord{
			x: c.x,
			y: c.y + 1,
		}
	case West:
		return Coord{
			x: c.x - 1,
			y: c.y,
		}
	case East:
		return Coord{
			x: c.x + 1,
			y: c.y,
		}
	default:
		panic("unknown direction")
	}
}
