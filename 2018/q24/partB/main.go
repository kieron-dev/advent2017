package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q24"
)

func main() {
	for i := 1; i < 100; i++ {
		if ImmuneWins(i) {
			break
		}
	}
}

func ImmuneWins(boost int) bool {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	s := q24.NewSystem(f)
	s.BoostImmuneSystem(boost)
	s.EliminateEnemy()
	if s.ImmuneUnits() > 0 && s.InfectionUnits() == 0 {
		fmt.Printf("s.ImmuneUnits() = %+v\n", s.ImmuneUnits())
		return true
	}
	return false
}
