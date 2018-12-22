package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q18"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	a := q18.NewArea(f)
	score := a.GetBigFutureScore(1000000000)
	a.Print()
	fmt.Printf("score = %+v\n", score)
}
