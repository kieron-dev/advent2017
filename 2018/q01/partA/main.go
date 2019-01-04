package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q01"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c := q01.NewCalibrator(f)
	sum := c.Add()

	fmt.Printf("sum = %+v\n", sum)
}
