package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q14"
)

func main() {
	recipes := q14.NewRecipes()
	scores := recipes.ScoresAfter(765071)
	fmt.Printf("scores = %+v\n", scores)
}
