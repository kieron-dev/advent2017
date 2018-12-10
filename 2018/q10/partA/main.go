package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q10"
)

func main() {
	f := q10.NewField()
	f.Load(os.Stdin)

	for i := 1; i < 100000; i++ {
		f.Step()
		img := f.MakePNG()
		if img == nil {
			continue
		}
		fmt.Printf("i = %+v\n", i)
		f.PrintAscii()
		out, err := os.Create(fmt.Sprintf("%05d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		png.Encode(out, img)
	}
}
