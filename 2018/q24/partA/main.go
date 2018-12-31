package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q24"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	s := q24.NewSystem(f)
	s.EliminateEnemy()
	fmt.Printf("s.ImmuneUnits() = %+v\n", s.ImmuneUnits())
	fmt.Printf("s.InfectionUnits() = %+v\n", s.InfectionUnits())
}
