package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q09"
)

func main() {
	game := q09.NewGame(476, 7143100)
	fmt.Println(game.Play())
}
