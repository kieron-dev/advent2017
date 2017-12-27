package particles

type Vector struct {
	x float64
	y float64
	z float64
}

func (v Vector) ManhattanDistance() float64 {
	return v.x + v.y + v.z
}

func (v Vector) Add(u Vector) Vector {
	return NewVector(v.x+u.x, v.y+u.y, v.z+u.z)
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

type Particle struct {
	position     Vector
	velocity     Vector
	acceleration Vector
	id           int
}

func New(pos, vel, acc Vector, id int) Particle {
	return Particle{
		position:     pos,
		velocity:     vel,
		acceleration: acc,
		id:           id,
	}
}

func (p Particle) ManhattanDistance() float64 {
	return p.position.ManhattanDistance()
}

func (p Particle) Position(t int) Vector {
	for i := 0; i < t; i++ {
		p.velocity = p.velocity.Add(p.acceleration)
		p.position = p.position.Add(p.velocity)
	}
	return p.position
}
