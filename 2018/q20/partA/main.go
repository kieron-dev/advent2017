package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q20"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	p := q20.NewPlan(f)
	p.ProcessRegex()
	res, count := p.FurthestRoom()
	fmt.Printf("res = %+v\n", res)
	fmt.Printf("count = %+v\n", count)
}
