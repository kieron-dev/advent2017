package main

import (
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q12"
)

func main() {
	plants := q12.NewPlants(os.Stdin)
	for i := 0; i < 20; i++ {
		plants.Step()
	}
	fmt.Println(plants.HashPosSum())
}
