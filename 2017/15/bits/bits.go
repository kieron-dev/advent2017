package bits

const (
	factorA = 16807
	factorB = 48271
)

func GetNext(prev, factor int64) int64 {
	return (prev * factor) % 2147483647
}

func GetNextA(prev int64) int64 {
	return GetNext(prev, factorA)
}

func GetNextB(prev int64) int64 {
	return GetNext(prev, factorB)
}

func GetNextA4(prev int64) int64 {
	a := GetNext(prev, factorA)
	for a%4 != 0 {
		a = GetNext(a, factorA)
	}
	return a
}

func GetNextB8(prev int64) int64 {
	b := GetNext(prev, factorB)
	for b%8 != 0 {
		b = GetNext(b, factorB)
	}
	return b
}

func GetLow16Bits(num int64) int64 {
	mask := (int64(1) << 16) - 1
	return num & mask
}

func CountLowerMatches(startA, startB int64,
	aFunc, bFunc func(int64) int64,
	lim int) int {

	c := 0
	a := startA
	b := startB
	for i := 0; i < lim; i++ {
		a = aFunc(a)
		b = bFunc(b)
		if GetLow16Bits(a) == GetLow16Bits(b) {
			c++
		}
	}
	return c
}
