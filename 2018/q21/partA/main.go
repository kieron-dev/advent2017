package main

import "fmt"

const r3const = 4921097

func main() {
	r1 := 1 << 16
	r3 := r3const

	for {
		r3 += r1 & 255
		r3 &= (1 << 24) - 1
		r3 *= 65899
		r3 &= (1 << 24) - 1

		if r1 < 256 {
			fmt.Printf("r3 = %+v\n", r3)
			return
			// r1 = r3 | 1<<16
			// r3 = r3const
		} else {
			r1 /= 256
		}
	}
}
