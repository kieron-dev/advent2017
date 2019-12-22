package fft

import (
	"strconv"
)

type Transform struct {
}

func NewTransform() *Transform {
	return &Transform{}
}

func (t *Transform) Process(in []int) []int {
	out := []int{}
	for i := 0; i < len(in); i++ {
		n := 0
		for j := 0; j < len(in); j++ {
			n += Coeff(i, j) * in[j]
		}
		out = append(out, mod(n)%10)
	}
	return out
}

func (t *Transform) ProcessForOffset(in []int) []int {
	sum := 0
	for i := 0; i < len(in); i++ {
		sum = (sum + in[i]) % 10
	}

	out := make([]int, len(in))
	out[0] = sum
	for i := 1; i < len(out); i++ {
		out[i] = (10 + out[i-1] - in[i-1]) % 10
	}

	return out
}

func Coeff(out, inPos int) int {
	n := (inPos + 1) / (out + 1)
	switch n % 4 {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 0
	case 3:
		return -1
	default:
		panic("eh?")
	}
}

func mod(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func StringToSlice(s string) []int {
	out := make([]int, len(s))
	for i, r := range s {
		var err error
		out[i], err = strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
	}
	return out
}

func GetOffset(in string) int {
	out, err := strconv.Atoi(in[:7])
	if err != nil {
		panic(err)
	}

	return out
}
