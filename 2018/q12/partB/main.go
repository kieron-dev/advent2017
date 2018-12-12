package main

import (
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q12"
)

func main() {
	plants := q12.NewPlants(os.Stdin)
	for i := 0; i < 10000; i++ {
		plants.Step()
	}
	sum1 := plants.HashPosSum()

	for i := 0; i < 10000; i++ {
		plants.Step()
	}
	sum2 := plants.HashPosSum()

	fmt.Printf("sum1 = %+v\n", sum1)
	fmt.Printf("sum2 = %+v\n", sum2)

	sum := 49999999*(sum2-sum1) + sum1
	fmt.Printf("answer = %+v\n", sum)
}
