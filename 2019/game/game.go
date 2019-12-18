package game

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
	"github.com/nsf/termbox-go"
)

type TileType int

const (
	Empty TileType = iota
	Wall
	Block
	HPaddle
	Ball
)

var graphics = map[TileType]rune{
	Empty:   ' ',
	Wall:    '|',
	Block:   '#',
	HPaddle: '^',
	Ball:    'o',
}

type Game struct {
	computer       *intcode.Computer
	in, out        chan int
	tiles          map[grid.Coord]TileType
	minX, minY     int
	maxX, maxY     int
	mutex          sync.Mutex
	cells          []termbox.Cell
	tw, th         int
	ballX, paddleX int
	ballDir        int
}

func NewGame() *Game {
	g := Game{}

	g.in = make(chan int, 2)
	g.out = make(chan int, 3)
	g.computer = intcode.NewComputer(g.in, g.out)
	g.tiles = map[grid.Coord]TileType{}

	g.minX = 1000
	g.minY = 1000
	g.maxX = 0
	g.maxY = 0

	return &g
}

func (g *Game) Pay() {
	g.computer.SetAddr(0, 2)
}

func (g *Game) LoadProgram(prog io.Reader) {
	bytes, err := ioutil.ReadAll(prog)
	if err != nil {
		panic(err)
	}

	g.computer.SetInput(strings.TrimSpace(string(bytes)))
}

func (g *Game) SetJoystick(n int) {
	g.in <- n
}

func (g *Game) reallocTermBuffer(w, h int) {
	g.cells = make([]termbox.Cell, w*h)
	g.tw = w
	g.th = h
}

func (g *Game) Run() int {

	if os.Getenv("SHOW_GRID") == "true" {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
		g.reallocTermBuffer(termbox.Size())
	}

	if os.Getenv("MANUAL_CONTROL") == "true" {
		go func() {
			for {
				switch ev := termbox.PollEvent(); ev.Type {
				case termbox.EventKey:
					if ev.Key == termbox.KeyArrowLeft {
						g.in <- -1
					} else if ev.Key == termbox.KeyArrowDown {
						g.in <- 0
					} else if ev.Key == termbox.KeyArrowRight {
						g.in <- 1
					}
				}
			}
		}()
	}

	go func() {
		defer close(g.out)

		g.computer.Calculate()
	}()

	score := 0
	for x := range g.out {
		y := <-g.out
		tileTypeN := <-g.out

		tileType := TileType(tileTypeN)

		if x == -1 && y == 0 {
			score = tileTypeN
			continue
		}

		if tileType == Ball {
			if g.ballX != 0 {
				if g.ballX > x {
					g.ballDir = -1
				} else {
					g.ballDir = 1
				}
			}
			g.ballX = x

			if os.Getenv("MANUAL_CONTROL") != "true" {
				if g.ballDir == 1 && g.ballX-g.paddleX > 0 {
					g.in <- 1
				} else if g.ballDir == -1 && g.ballX-g.paddleX < 0 {
					g.in <- -1
				} else {
					g.in <- 0
				}
			}
		}

		if TileType(tileType) == HPaddle {
			g.paddleX = x
		}

		g.AddTile(int(x), int(y), TileType(tileType))
		if os.Getenv("SHOW_GRID") == "true" {
			g.updateFrame()
		}

	}
	return score
}

func (g *Game) updateFrame() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	copy(termbox.CellBuffer(), g.cells)
	for k, v := range g.tiles {
		termbox.SetCell(k.X(), k.Y(), graphics[v], termbox.ColorWhite, termbox.ColorDefault)
	}
	err = termbox.Flush()
	if err != nil {
		panic(err)
	}

}

func (g *Game) AddTile(x, y int, tileType TileType) {
	if x > g.maxX {
		g.maxX = x
	}
	if x < g.minX {
		g.minX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
	if y < g.minY {
		g.minY = y
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
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
