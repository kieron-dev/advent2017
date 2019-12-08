package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	image := advent2019.NewImage(6, 25)
	image.Load(strings.TrimSpace(string(bytes)))

	layer := image.Decode()
	layer.Print()
}
