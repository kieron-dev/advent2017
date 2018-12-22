package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q17"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	s := q17.NewSlice(f)
	s.Flow(q17.NewCoord(500, 0))
	s.Print()
	fmt.Printf("s.CountStaticWater() = %+v\n", s.CountStaticWater())
}
