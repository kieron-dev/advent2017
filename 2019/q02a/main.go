package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	all, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	c := advent2019.NewComputer(nil, nil)
	c.SetInput(strings.TrimSpace(string(all)))
	c.Prime(12, 02)
	out := c.Calculate()

	fmt.Printf("out = %+v\n", out)
}
