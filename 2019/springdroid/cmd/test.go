package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/springdroid"
)

func main() {
	d := springdroid.NewDroid()
	input, err := ioutil.ReadFile("days/input21")
	if err != nil {
		panic(err)
	}
	d.LoadProgram(strings.TrimSpace(string(input)))
	d.RunProgram()

	go func() {
		for {
			b := d.Output()
			fmt.Printf("%c", b)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		d.Input(line)
	}
}
