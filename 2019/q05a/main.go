package main

import (
	"bytes"
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

	in := bytes.NewBuffer([]byte("1\n"))

	c := advent2019.NewComputer(in)
	c.SetInput(strings.TrimSpace(string(all)))
	c.Calculate()
}
