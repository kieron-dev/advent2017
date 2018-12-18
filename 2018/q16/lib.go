package q16

type Computer struct {
	Registers [4]int
}

func NewComputer() *Computer {
	c := Computer{}
	return &c
}

func (c *Computer) Equals(other *Computer) bool {
	for i, r := range c.Registers {
		if r != other.Registers[i] {
			return false
		}
	}
	return true
}

func (comp *Computer) SetRegisters(a, b, c, d int) {
	comp.Registers[0] = a
	comp.Registers[1] = b
	comp.Registers[2] = c
	comp.Registers[3] = d
}

type Op func(a, b, c int)

func (c *Computer) Ops() []Op {
	return []Op{
		c.Addr,
		c.Addi,
		c.Mulr,
		c.Muli,
		c.Banr,
		c.Bani,
		c.Borr,
		c.Bori,
		c.Setr,
		c.Seti,
		c.Gtir,
		c.Gtri,
		c.Gtrr,
		c.Eqir,
		c.Eqri,
		c.Eqrr,
	}
}

func (c *Computer) Addr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] + c.Registers[R2]
}

func (c *Computer) Addi(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] + V2
}

func (c *Computer) Mulr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] * c.Registers[R2]
}

func (c *Computer) Muli(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] * V2
}

func (c *Computer) Banr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] & c.Registers[R2]
}

func (c *Computer) Bani(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] & V2
}

func (c *Computer) Borr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] | c.Registers[R2]
}

func (c *Computer) Bori(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] | V2
}

func (c *Computer) Setr(R1, _, R3 int) {
	c.Registers[R3] = c.Registers[R1]
}

func (c *Computer) Seti(V1, _, R3 int) {
	c.Registers[R3] = V1
}

func (c *Computer) Gtir(V1, R2, R3 int) {
	if V1 > c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Gtri(R1, V2, R3 int) {
	if c.Registers[R1] > V2 {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Gtrr(R1, R2, R3 int) {
	if c.Registers[R1] > c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqir(V1, R2, R3 int) {
	if V1 == c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqri(R1, V2, R3 int) {
	if c.Registers[R1] == V2 {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqrr(R1, R2, R3 int) {
	if c.Registers[R1] == c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) MatchingOps(d *Computer, x, y, z int) []int {
	matches := []int{}
	var savedRegs [4]int
	for i := 0; i < 4; i++ {
		savedRegs[i] = c.Registers[i]
	}

	for n, op := range c.Ops() {
		for i := 0; i < 4; i++ {
			c.Registers[i] = savedRegs[i]
		}
		op(x, y, z)
		if c.Equals(d) {
			matches = append(matches, n)
		}
	}
	return matches
}
