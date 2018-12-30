package main

import "fmt"

func main() {
	mn := 0
	mr0 := 0
	lim := 1 << 24
	for r0 := 0; r0 < lim; r0++ {
		if r0%100000 == 0 {
			fmt.Printf("...%d\n", r0)
		}
		n := doIt(r0)
		if n > 0 && n > mn {
			mn = n
			mr0 = r0
			fmt.Printf("mn = %+v\n", mn)
			fmt.Printf("mr0 = %+v\n", mr0)
		}
	}
	fmt.Println("-------")
	fmt.Printf("mn = %+v\n", mn)
	fmt.Printf("mr0 = %+v\n", mr0)

}

const r3const = 4921097

func doIt(r0 int) int {
	r1 := 1 << 16
	r3 := r3const
	inst := 0

	i := 0
	for {
		i++
		if i > 50000 {
			return -1
		}
		r3 += r1 & 255
		r3 &= (1 << 24) - 1
		r3 *= 65899
		r3 &= (1 << 24) - 1

		if r1 < 256 {
			if r3 == r0 {
				return inst
			}
			r1 = r3 | 1<<16
			r3 = r3const
		} else {
			r1 /= 256
			inst += r1
		}
	}
}
