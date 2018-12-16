package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q14"
)

func main() {
	r := q14.NewRecipes()
	res := r.ScoresBefore("765071")
	fmt.Printf("res = %+v\n", res)
}
