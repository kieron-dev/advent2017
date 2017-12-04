package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/02/checksum"
)

func main() {
	usage := fmt.Sprintf("%s <captchaFile>", os.Args[0])
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

	arr := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "\t")
		numSlice := []int{}
		for _, nstr := range nums {
			val, _ := strconv.Atoi(nstr)
			numSlice = append(numSlice, val)
		}
		arr = append(arr, numSlice)
	}

	fmt.Println(checksum.Calc(arr))
}
