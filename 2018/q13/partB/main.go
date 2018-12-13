package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q13"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	mine := q13.NewMine(f)
	i, cart := mine.RunTillOneLeft()

	fmt.Printf("i = %+v\n", i)
	fmt.Printf("cart = %+v\n", cart)
}
