package game

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
	"github.com/kieron-pivotal/advent2017/advent2019/grid"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Block
	HPaddle
	Ball
)

type Game struct {
	computer *advent2019.Computer
	in, out  chan string
	tiles    map[grid.Coord]TileType
}

func NewGame() *Game {
	g := Game{}

	g.in = make(chan string)
	g.out = make(chan string)
	g.computer = advent2019.NewComputer(g.in, g.out)
	g.tiles = map[grid.Coord]TileType{}

	return &g
}

func (g *Game) LoadProgram(prog io.Reader) {
	bytes, err := ioutil.ReadAll(prog)
	if err != nil {
		panic(err)
	}

	g.computer.SetInput(strings.TrimSpace(string(bytes)))
}

func (g *Game) Run() {
	go func() {
		defer close(g.out)

		g.computer.Calculate()
	}()

	for xStr := range g.out {
		yStr := <-g.out
		tileTypeStr := <-g.out

		x, err := strconv.Atoi(xStr)
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(yStr)
		if err != nil {
			panic(err)
		}
		tileType, err := strconv.Atoi(tileTypeStr)
		if err != nil {
			panic(err)
		}

		g.AddTile(x, y, TileType(tileType))
	}
}

func (g *Game) AddTile(x, y int, tileType TileType) {
	g.tiles[grid.NewCoord(x, y)] = tileType
}

func (g *Game) TileCount(t TileType) int {
	count := 0
	for _, v := range g.tiles {
		if v != t {
			continue
		}
		count++
	}
	return count
}
