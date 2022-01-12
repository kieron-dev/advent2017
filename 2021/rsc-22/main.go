package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	text := strings.Fields(r.Replace(string(data)))

	type dim struct {
		min, max int
	}
	type flip struct {
		on int
		p  [3]dim
	}

	var prog []*flip

	const Off = 130_000

	for ; len(text) > 0; text = text[7:] {
		var f flip
		f.on,
			f.p[0].min, f.p[0].max,
			f.p[1].min, f.p[1].max,
			f.p[2].min, f.p[2].max =
			atoi(text[0]),
			atoi(text[1])+Off,
			atoi(text[2])+Off+1,
			atoi(text[3])+Off,
			atoi(text[4])+Off+1,
			atoi(text[5])+Off,
			atoi(text[6])+Off+1

		prog = append(prog, &f)
	}
	const N = 850

	var remap [3][2 * Off]int
	var width [3][N]int
	for _, f := range prog {
		for i, d := range f.p {
			remap[i][d.min] = 1
			remap[i][d.max] = 1
		}
	}

	for i := range remap {
		t := 0
		for j, v := range remap[i] {
			t += v
			remap[i][j] = t
			width[i][t]++
		}
	}

	for _, f := range prog {
		for i := range f.p {
			f.p[i].min = remap[i][f.p[i].min]
			f.p[i].max = remap[i][f.p[i].max]
		}
	}

	var sw [N][N][N]byte

	for _, f := range prog {
		for i := f.p[0].min; i < f.p[0].max; i++ {
			for j := f.p[1].min; j < f.p[1].max; j++ {
				for k := f.p[2].min; k < f.p[2].max; k++ {
					sw[i][j][k] = byte(f.on)
				}
			}
		}
	}

	total := 0
	for i := range sw {
		for j := range sw {
			for k := range sw {
				total += width[0][i] * width[1][j] * width[2][k] * int(sw[i][j][k])
			}
		}
	}

	fmt.Println(total)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

var r = strings.NewReplacer(
	"x=", "",
	"y=", "",
	"z=", "",
	"on", "1",
	"off", "0",
	",", " ",
	"..", " ",
)
