package main

import "fmt"

func main() {
	mn := 270
	mr0 := 0
	for r0 := 0; r0 < 256000000; r0++ {
		n := doIt(r0, mn)
		if n > 0 && n < mn {
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

func doIt(r0, mn int) int {
	r1 := 1 << 16
	r3 := r3const
	inst := 0

	for {
		inst++
		if inst > mn {
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
