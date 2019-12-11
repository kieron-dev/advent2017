package grid

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
