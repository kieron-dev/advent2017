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
	fight := q15.NewFight(f)
	res := fight.Run()
	fmt.Printf("res = %+v\n", res)
}
