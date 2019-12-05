package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	gen := advent2019.NewPWGen(347312, 805915, true)
	i := 1
	for {
		_, err := gen.Next()
		if err != nil {
			break
		}
		i++
	}
	fmt.Printf("res = %d\n", i)
}
