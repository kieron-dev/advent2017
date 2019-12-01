package advent2019

type FuelCalc struct{}

func (c FuelCalc) FuelForMass(m int) int {
	return m/3 - 2
}

func (c FuelCalc) FuelForFuel(m int) int {
	total := 0
	for {
		extra := c.FuelForMass(m)
		if extra < 1 {
			break
		}
		total += extra
		m = extra
	}
	return total
}
