package main

import "fmt"

const r3const = 4921097

func main() {
	haltingNs := []int{}
	haltingSet := map[int]bool{}

	r1 := 1 << 16
	r3 := r3const

	for {
		r3 += r1 & 255
		r3 &= (1 << 24) - 1
		r3 *= 65899
		r3 &= (1 << 24) - 1

		if r1 < 256 {
			if haltingSet[r3] {
				break
			}
			haltingSet[r3] = true
			haltingNs = append(haltingNs, r3)

			r1 = r3 | 1<<16
			r3 = r3const
		} else {
			r1 /= 256
		}
	}

	res := haltingNs[len(haltingNs)-1]
	fmt.Printf("res = %+v\n", res)
}
