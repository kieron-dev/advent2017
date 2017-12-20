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
	m1 := assembly.NewMachine(0)
	m2 := assembly.NewMachine(1)
	m1.Duet(m2)

	for scanner.Scan() {
		instr := scanner.Text()
		m1.AppendInstruction(instr)
		m2.AppendInstruction(instr)
	}

	assembly.RunMachines([]*assembly.Machine{m1, m2})

	fmt.Println("Part2:", m2.GetCount())

}
