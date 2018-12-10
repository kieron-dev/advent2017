package q10

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"io"
)

type Point struct {
	x  int
	y  int
	vx int
	vy int
}

func (p *Point) Step() {
	p.x += p.vx
	p.y += p.vy
}

type Field struct {
	points []*Point
}

func NewField() *Field {
	f := Field{}
	f.points = []*Point{}
	return &f
}

func (f *Field) Load(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		p := Point{}
		fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &p.x, &p.y, &p.vx, &p.vy)
		f.points = append(f.points, &p)
	}
}

func (f *Field) Step() {
	for _, p := range f.points {
		p.Step()
	}
}

func (f *Field) Boundaries() (minX, minY, maxX, maxY int) {
	minX = f.points[0].x
	maxX = minX
	minY = f.points[0].y
	maxY = minY

	for i := 1; i < len(f.points); i++ {
		p := f.points[i]
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return
}

func (f *Field) MakeGrid() [][]bool {
	minX, minY, maxX, maxY := f.Boundaries()
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]bool, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]bool, width)
	}
	for _, p := range f.points {
		grid[p.y-minY][p.x-minX] = true
	}

	return grid
}

func (f *Field) MakePNG() *image.RGBA {
	minX, minY, maxX, maxY := f.Boundaries()
	w := maxX - minX + 1
	h := maxY - minY + 1
	if h > 10 {
		return nil
	}
	upLeft := image.Point{0, 0}
	lowRight := image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	for _, p := range f.points {
		img.Set(p.x-minX, p.y-minY, white)
	}
	return img
}

func (f *Field) PrintAscii() {
	grid := f.MakeGrid()
	if len(grid) > 100 || len(grid[0]) > 100 {
		return
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
