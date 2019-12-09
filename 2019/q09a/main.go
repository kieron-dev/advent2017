package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	progBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	in := make(chan big.Int, 1)
	out := make(chan big.Int, 20)

	c := advent2019.NewComputer(in, out)
	c.SetInput(strings.TrimSpace(string(progBytes)))
	in <- *big.NewInt(1)
	c.Calculate()

	close(out)
	for n := range out {
		fmt.Printf("n.String() = %+v\n", n.String())
	}
}
