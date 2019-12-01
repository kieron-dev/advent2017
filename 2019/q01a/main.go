package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	fuel := 0
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	fileReader := advent2019.FileReader{}
	fuelCalc := advent2019.FuelCalc{}

	fileReader.Each(file, func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		fuel += fuelCalc.FuelForMass(n)
	})

	fmt.Printf("fuel required: %d\n", fuel)

}
