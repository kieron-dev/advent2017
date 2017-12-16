package defrag

import (
	"encoding/hex"
	"fmt"

	"github.com/kieron-pivotal/advent2017/10/hash"
)

func Hash(seed string, row int) string {
	in := fmt.Sprintf("%s-%d", seed, row)
	return hash.Compute2([]byte(in), 256)
}

func BitCount(hexIn string) int {
	bs, _ := hex.DecodeString(hexIn)
	t := 0
	for _, b := range bs {
		m := byte(1)
		for i := 0; i < 8; i++ {
			// fmt.Println(m, b, m&b)
			if m&b > 0 {
				t++
			}
			m <<= 1
		}
	}
	return t
}

func CountUsed(in string) int {
	c := 0
	for i := 0; i < 128; i++ {
		c += BitCount(Hash(in, i))
	}
	return c
}
