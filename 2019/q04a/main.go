package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	gen := advent2019.NewPWGen(347312, 805915, false)
	var err error
	i := 0
	for err == nil {
		_, err = gen.Next()
		i++
	}
	fmt.Printf("res = %d\n", i)
}
