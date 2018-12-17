package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q15"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	res := q15.RunWithNoElfDeath(f)
	fmt.Printf("res = %+v\n", res)
}
