package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	in := make(chan int, 10)
	out := make(chan int, 10)

	in <- 5
	c := advent2019.NewComputer(in, out)
	c.SetInput(strings.TrimSpace(string(all)))
	c.Calculate()

	res := <-out
	fmt.Printf("res = %+v\n", res)
}
