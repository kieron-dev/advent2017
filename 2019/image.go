package advent2019

import (
	"fmt"
	"math"
	"strconv"
)

type Image struct {
	layers []*Layer
	height int
	width  int
}

func NewImage(height, width int) *Image {
	i := Image{
		height: height,
		width:  width,
		layers: []*Layer{},
	}
	return &i
}

func (i *Image) Load(pixels string) {
	for l := 0; l*i.width*i.height < len(pixels); l++ {
		layer := Layer{pixels: [][]int{}}
		for y := 0; y < i.height; y++ {
			row := []int{}
			for x := 0; x < i.width; x++ {
				pstr := pixels[l*i.width*i.height+y*i.width+x]
				n, err := strconv.Atoi(string(pstr))
				if err != nil {
					panic(err)
				}
				row = append(row, n)
			}
			layer.pixels = append(layer.pixels, row)
		}
		i.layers = append(i.layers, &layer)
	}
}

func (i *Image) GetLayer(idx int) *Layer {
	return i.layers[idx]
}

func (i *Image) Decode() *Layer {
	layer := &Layer{pixels: [][]int{}}

	for y := 0; y < i.height; y++ {
		row := []int{}
		for x := 0; x < i.width; x++ {
			var p int
			for l := len(i.layers) - 1; l >= 0; l-- {
				px := i.layers[l].pixels[y][x]
				if px != 2 {
					p = px
				}
			}
			row = append(row, p)
		}
		layer.pixels = append(layer.pixels, row)
	}

	return layer
}

func (i *Image) FindLayerWithFewestZeros() *Layer {
	minZeros := math.MaxInt32
	var minLayer *Layer
	for _, layer := range i.layers {
		zeros := layer.Count(0)
		if zeros < minZeros {
			minZeros = zeros
			minLayer = layer
		}
	}
	return minLayer
}

type Layer struct {
	pixels [][]int
}

func (l *Layer) Count(n int) int {
	c := 0
	for _, r := range l.pixels {
		for _, p := range r {
			if p == n {
				c++
			}
		}
	}
	return c
}

func (l *Layer) Print() {
	for _, row := range l.pixels {
		for _, p := range row {
			if p == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}
