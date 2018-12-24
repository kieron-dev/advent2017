package main

import "fmt"

func main() {
	num := 10551383
	res := 0

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			res += i
		}
	}
	fmt.Printf("res = %+v\n", res)
}
