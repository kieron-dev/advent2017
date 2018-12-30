package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q23"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	t := q23.NewTeleport(f)
	strongest := t.Strongest()
	InRange := t.InRange(strongest)
	fmt.Printf("InRange = %+v\n", InRange)
}
