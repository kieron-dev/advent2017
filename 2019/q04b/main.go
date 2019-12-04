package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	gen := advent2019.NewPWGen(347312, 805915, true)
	i := 1
	fmt.Printf("gen.Current() = %+v\n", gen.Current())
	for {
		_, err := gen.Next()
		if err != nil {
			break
		}
		fmt.Printf("gen.Current() = %+v\n", gen.Current())
		i++
	}
	fmt.Printf("res = %d\n", i)
}
