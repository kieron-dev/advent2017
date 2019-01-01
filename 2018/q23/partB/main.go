package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q23"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	t := q23.NewTeleport(f)
	bestCoord := t.FindBestCoord(20)
	fmt.Printf("bestCoord = %+v\n", bestCoord)
	fmt.Printf("dist = %+v\n", q23.Abs(bestCoord.X)+q23.Abs(bestCoord.Y)+q23.Abs(bestCoord.Z))

}
