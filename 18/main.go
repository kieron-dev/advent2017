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
	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)
	m1 := assembly.NewMachine(0, ch2, ch1)
	m2 := assembly.NewMachine(1, ch1, ch2)

	for scanner.Scan() {
		instr := scanner.Text()
		m1.AppendInstruction(instr)
		m2.AppendInstruction(instr)
	}

	assembly.RunMachines([]*assembly.Machine{m1, m2})

	fmt.Println("Part2:", m2.GetCount())

	// fmt.Println("Part1:", m.RecoverVal())
	// fmt.Println("Part2:", firewall.MinDelay(config))
}
