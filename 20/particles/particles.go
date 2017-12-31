package particles

import (
	"sort"
)

type Separation struct {
	Id           int
	DistFromLast int
}

type Distribution struct {
	Particles []Separation
}

func (d *Distribution) HasSteadyOrder(d0 *Distribution) bool {
	if len(d.Particles) != len(d0.Particles) {
		return false
	}
	for i, p := range d.Particles {
		p0 := d0.Particles[i]
		if p0.Id != p.Id {
			return false
		}
		if p0.DistFromLast > p.DistFromLast {
			return false
		}
	}
	return true
}

func DistrFromParticles(ps []*Particle) *Distribution {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].distance < ps[j].distance
	})
	prevDist := 0
	d := Distribution{Particles: []Separation{}}
	for _, p := range ps {
		d.Particles = append(d.Particles,
			Separation{
				Id:           p.id,
				DistFromLast: p.distance - prevDist,
			})
		prevDist = p.distance
	}
	return &d
}

type Vector struct {
	x int
	y int
	z int
}

func (v Vector) ManhattanDistance() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (v Vector) Add(u Vector) Vector {
	return NewVector(v.x+u.x, v.y+u.y, v.z+u.z)
}

func NewVector(x, y, z int) Vector {
	return Vector{x, y, z}
}

type Particle struct {
	position     Vector
	velocity     Vector
	acceleration Vector
	id           int
	distance     int
	time         int
}

func New(pos, vel, acc Vector, id int) *Particle {
	p := Particle{
		position:     pos,
		velocity:     vel,
		acceleration: acc,
		id:           id,
	}
	p.distance = p.position.ManhattanDistance()
	return &p
}

func (p *Particle) Advance(t int) {
	for i := 0; i < t; i++ {
		p.velocity = p.velocity.Add(p.acceleration)
		p.position = p.position.Add(p.velocity)
	}
	p.distance = p.position.ManhattanDistance()
	p.time += t
}

func (p *Particle) Position() Vector {
	return p.position
}

func (p *Particle) Distance() int {
	return p.distance
}

func (p *Particle) Time() int {
	return p.time
}

func (p *Particle) Id() int {
	return p.id
}

func GetEventualClosest(ps []*Particle) *Particle {
	dLast := DistrFromParticles(ps)
	steadyCount := 0
	requiredSteadyCount := 2
	for {
		for _, p := range ps {
			p.Advance(1)
		}
		dNext := DistrFromParticles(ps)
		if dNext.HasSteadyOrder(dLast) {
			steadyCount++
			if steadyCount == requiredSteadyCount {
				return ps[0]
			}
		} else {
			steadyCount = 0
		}
		dLast = dNext
	}
}

func RemoveCollisions(ps []*Particle) []*Particle {
	ret := []*Particle{}
	positions := map[Vector][]*Particle{}
	for _, p := range ps {
		positions[p.position] = append(positions[p.position], p)
	}
	for _, particles := range positions {
		if len(particles) == 1 {
			ret = append(ret, particles[0])
		}
	}
	return ret
}

func GetRemainingAfterAllCollisions(ps []*Particle, iterations int) int {
	for i := 0; i < iterations; i++ {
		ps = RemoveCollisions(ps)
		for _, p := range ps {
			p.Advance(1)
		}
	}
	return len(ps)
}
