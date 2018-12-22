package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kieron-pivotal/advent2017/2018/q18"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	a := q18.NewArea(f)
	for {
		fmt.Println("")
		a.Print()
		a.Step()
		time.Sleep(50 * time.Millisecond)
	}
}
