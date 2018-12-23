package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q19"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c := q19.NewComputer(f)
	res := c.Execute()
	fmt.Printf("res = %+v\n", res)
}
