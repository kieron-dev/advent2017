package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	target := 19690720

	all, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			c := advent2019.NewComputer(nil, nil)
			c.SetInput(strings.TrimSpace(string(all)))
			c.Prime(noun, verb)

			if c.TryCalculate() == target {
				fmt.Printf("%d\n", 100*noun+verb)
				return
			}
		}
	}

}
