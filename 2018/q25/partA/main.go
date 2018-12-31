package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q25"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	s := q25.NewSpace(f)
	p := s.Partition()
	fmt.Printf("p = %+v\n", p)
}
