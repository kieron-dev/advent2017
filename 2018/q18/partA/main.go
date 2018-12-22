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
	for i := 0; i < 10; i++ {
		a.Step()
	}
	fmt.Printf("a.Score() = %+v\n", a.Score())
}
