package main

func main() {
	var (
		r0 int
		r1 int
		r2 int
		r3 int
		r4 int
		// r5 int
	)

	r0 = 1

	//0
	// start at 17

	//1
	r2 = 1

	//2
	r4 = 1

	//4
	if r2*r4 == r1 {
		//7
		r0 += r2
	}

	//8
	r4 += 1

	//9
	if r4 <= r1 {
		// goto 3
	}

	//12
	r2 += 1

	//13
	if r2 > r1 {
		r3 = 1
		// exit
	} else {
		r3 = 0
		// goto 2
	}

	//17
	r1 += 2

	//18
	r1 *= r1

	//19
	r1 *= 19

	//20
	r1 *= 11

	//21
	r3 += 6

	//22
	r3 *= 22

	//23
	r3 += 15

	//24
	r1 += r3

	//25
	// skip r0 (1), i.e. goto 27

	//26
	// goto 1

	//27
	r3 = 27

	//28
	r3 *= 28

	//29
	r3 += 29

	//30
	r3 *= 30

	//31
	r3 *= 14

	//32
	r3 *= 32

	//33
	r1 += r3

	//34
	r0 = 0

	//35
	// goto 1

}
