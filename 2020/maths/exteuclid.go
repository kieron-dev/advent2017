package maths

import "math/big"

// ExtEuclid returns a triple (d,s,t) such that d = gcd(a,b) and
// d == a*s + b*t
func ExtEuclid(a, b *big.Int) (d, s, t *big.Int) {
	if b.Int64() == 0 {
		return a, big.NewInt(1), big.NewInt(0)
	}

	aModB := new(big.Int).Mod(a, b)
	d1, s1, t1 := ExtEuclid(b, aModB)

	d = d1
	s = t1

	aDivB := new(big.Int).Div(a, b)
	timesT1 := new(big.Int).Mul(aDivB, t1)
	t = new(big.Int).Sub(s1, timesT1)

	return
}
