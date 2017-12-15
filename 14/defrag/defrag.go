package defrag

import (
	"encoding/hex"
	"fmt"

	"github.com/kieron-pivotal/advent2017/10/hash"
	"github.com/kieron-pivotal/advent2017/12/graph"
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

func HexToBin(in string) string {
	out := ""
	bs, _ := hex.DecodeString(in)
	for _, b := range bs {
		out += fmt.Sprintf("%08b", b)
	}
	return out
}

func HashesToForest(hashes []string) *graph.Graph {
	g := graph.New()
	for n := 0; n < 128*128; n++ {
		r := n / 128
		c := n % 128
		if hashes[r][c] == '0' {
			continue
		}
		children := []int{}
		if c < 127 && hashes[r][c+1] == '1' {
			children = append(children, n+1)
		}
		if r < 127 && hashes[r+1][c] == '1' {
			children = append(children, n+128)
		}
		g.LinkNodes(n, children)
	}
	return g
}

func CountBlocks(in string) int {
	hashes := []string{}
	for i := 0; i < 128; i++ {
		hashes = append(hashes, HexToBin(Hash(in, i)))
	}
	g := HashesToForest(hashes)
	return g.Groups()
}
