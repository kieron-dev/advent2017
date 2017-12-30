package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/kieron-pivotal/advent2017/20/particles"
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

	re := regexp.MustCompile("<(.*),(.*),(.*)>.*<(.*),(.*),(.*)>.*<(.*),(.*),(.*)>")

	id := 0
	ps := []*particles.Particle{}
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		components := make([]int, 9)
		for i := 0; i < 9; i++ {
			n, err := strconv.Atoi(match[i+1])
			if err != nil {
				panic(err)
			}
			components[i] = n
		}
		p := particles.New(
			particles.NewVector(components[0], components[1], components[2]),
			particles.NewVector(components[3], components[4], components[5]),
			particles.NewVector(components[6], components[7], components[8]),
			id,
		)
		id++
		ps = append(ps, p)
	}

	nearest := particles.GetEventualClosest(ps)
	fmt.Println("Part1:", nearest.Id(), nearest.Time())

}
