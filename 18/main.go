package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/18/assembly"
)

func main() {
	usage := fmt.Sprintf("%s <inputPath>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	m := assembly.NewMachine()

	for scanner.Scan() {
		instr := scanner.Text()
		m.AppendInstruction(instr)
	}

	m.Run()

	fmt.Println("Part1:", m.RecoverVal())
	// fmt.Println("Part2:", firewall.MinDelay(config))
}
