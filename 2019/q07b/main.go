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
	contents, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	max := big.NewInt(0)

	permHelper := advent2019.Perms{}
	for _, perm := range permHelper.All([]int64{5, 6, 7, 8, 9}) {
		arr := advent2019.NewFeedbackArray(5)
		arr.SetProgram(strings.TrimSpace(string(contents)))
		arr.SetPhase(perm)
		arr.WriteInitialInput(0)
		arr.Run()
		val := arr.GetResult()
		if val.Cmp(max) > 0 {
			max.Set(val)
		}
	}
	fmt.Printf("max = %+v\n", max.String())
}
