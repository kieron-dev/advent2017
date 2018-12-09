package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q09"
)

func main() {
	game := q09.NewGame(476, 71431)
	fmt.Println(game.Play())
}
