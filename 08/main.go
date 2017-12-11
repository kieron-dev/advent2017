package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/08/registers"
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

	registers := registers.Set{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		registers.AddInstruction(line)
	}

	registers.Process()

	fmt.Println(registers.GetMaxRegister())
}
