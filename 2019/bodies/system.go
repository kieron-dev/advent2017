package bodies

import (
	"fmt"
	"io"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

type Coord3 struct {
	coords [3]int
}

func NewCoord3(x, y, z int) Coord3 {
	c := Coord3{}
	c.coords = [3]int{x, y, z}
	return c
}

func (c Coord3) Plus(d Coord3) Coord3 {
	return Coord3{
		[3]int{
			c.coords[0] + d.coords[0],
			c.coords[1] + d.coords[1],
			c.coords[2] + d.coords[2],
		},
	}
}

func (c Coord3) X() int {
	return c.coords[0]
}

func (c Coord3) Y() int {
	return c.coords[1]
}

func (c Coord3) Z() int {
	return c.coords[2]
}

type Body struct {
	position Coord3
	velocity Coord3
}

func (b *Body) Pos() Coord3 {
	return b.position
}

func (b *Body) Vel() Coord3 {
	return b.velocity
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (b *Body) Energy() int {
	s1 := 0
	for i := 0; i < 3; i++ {
		s1 += abs(b.position.coords[i])
	}
	s2 := 0
	for i := 0; i < 3; i++ {
		s2 += abs(b.velocity.coords[i])
	}
	return s1 * s2
}

type System struct {
	moons         []*Body
	initialCoords []Coord3
}

func NewSystem() *System {
	s := System{}

	return &s
}

func (s *System) Load(state io.Reader) {
	r := advent2019.FileReader{}
	r.Each(state, func(line string) {
		moon := Body{velocity: Coord3{}}
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moon.position = NewCoord3(x, y, z)
		s.moons = append(s.moons, &moon)
		s.initialCoords = append(s.initialCoords, moon.position)
	})
}

func (s *System) Reset() {
	for i, c := range s.initialCoords {
		s.moons[i].position = c
		s.moons[i].velocity = NewCoord3(0, 0, 0)
	}
}

func (s *System) Moons() []*Body {
	return s.moons
}

func (s *System) Tick() {
	for m1 := 0; m1 < len(s.moons); m1++ {
		for m2 := m1 + 1; m2 < len(s.moons); m2++ {
			moon1 := s.moons[m1]
			moon2 := s.moons[m2]
			for i := 0; i < 3; i++ {
				if moon1.position.coords[i] < moon2.position.coords[i] {
					moon1.velocity.coords[i]++
					moon2.velocity.coords[i]--
				} else if moon1.position.coords[i] > moon2.position.coords[i] {
					moon1.velocity.coords[i]--
					moon2.velocity.coords[i]++
				}
			}
		}
	}
	for m := 0; m < len(s.moons); m++ {
		s.moons[m].position = s.moons[m].position.Plus(s.moons[m].velocity)
	}
}

func (s *System) TotalEnergy() int {
	t := 0
	for _, m := range s.moons {
		t += m.Energy()
	}
	return t
}

func (s *System) FirstXRepeat() int {
	s.Reset()
	return s.firstRepeat(0)
}

func (s *System) FirstYRepeat() int {
	s.Reset()
	return s.firstRepeat(1)
}

func (s *System) FirstZRepeat() int {
	s.Reset()
	return s.firstRepeat(2)
}

func (s *System) FirstRepeat() int64 {
	x := s.FirstXRepeat()
	y := s.FirstYRepeat()
	z := s.FirstZRepeat()

	return lcm(lcm(int64(x), int64(y)), int64(z))
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func (s *System) firstRepeat(idx int) int {
	initVals := []int{}
	for i := 0; i < len(s.moons); i++ {
		initVals = append(initVals, s.moons[i].position.coords[idx])
	}

	i := 0
	for {
		i++
		s.Tick()

		ok := true
		for j := 0; j < len(s.moons); j++ {
			m := s.moons[j]
			if m.position.coords[idx] != initVals[j] ||
				m.velocity.coords[idx] != 0 {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
}
