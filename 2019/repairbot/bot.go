package repairbot

import (
	"os"

	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/intcode"
	"github.com/nsf/termbox-go"
)

type StatusCode int

const (
	HitWall StatusCode = iota
	Success
	FoundOxygen
)

type Type int

const (
	Unknown Type = iota
	Wall
	Passable
	Oxygen
	Droid
)

var graphics = map[Type]rune{
	Unknown:  ' ',
	Wall:     '#',
	Passable: '.',
	Oxygen:   'o',
	Droid:    'D',
}

type Bot struct {
	in, out    chan int
	computer   *intcode.Computer
	cells      []termbox.Cell
	tw, th     int
	chart      map[grid.Coord]Type
	prevDirs   map[grid.Coord]grid.Direction
	position   grid.Coord
	oxygenPos  grid.Coord
	startPos   grid.Coord
	dists      map[grid.Coord]int
	oxygenDist int
}

func New() *Bot {
	b := Bot{
		in:       make(chan int),
		out:      make(chan int),
		position: grid.NewCoord(30, 30),
		chart:    map[grid.Coord]Type{},
		prevDirs: map[grid.Coord]grid.Direction{},
		dists:    map[grid.Coord]int{},
	}
	b.startPos = b.position
	b.dists[b.startPos] = 0
	b.computer = intcode.NewComputer(b.in, b.out)

	return &b
}

func (b *Bot) reallocTermBuffer(w, h int) {
	b.cells = make([]termbox.Cell, w*h)
	b.tw = w
	b.th = h
}

func (b *Bot) eventLoop(exit chan struct{}) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowLeft {
				b.Move(grid.West)
			} else if ev.Key == termbox.KeyArrowDown {
				b.Move(grid.South)
			} else if ev.Key == termbox.KeyArrowRight {
				b.Move(grid.East)
			} else if ev.Key == termbox.KeyArrowUp {
				b.Move(grid.North)
			} else if ev.Key == termbox.KeyCtrlQ {
				close(exit)
			}
		}
	}
}

func (b *Bot) exposeMap() {
	pos := b.position
	branchPoints := []grid.Coord{}

outer:
	for {
		options := []grid.Direction{}
		for _, d := range []grid.Direction{grid.North, grid.South, grid.East, grid.West} {
			// time.Sleep(time.Millisecond * 2)
			next := pos.Move(d)
			if _, alreadySet := b.chart[next]; alreadySet {
				continue
			}
			result := b.Move(d)
			switch result {
			case HitWall:
				continue
			case Success:
				options = append(options, d)
				b.Move(d.Opposite())
				b.prevDirs[next] = d.Opposite()
			case FoundOxygen:
				options = append(options, d)
				b.Move(d.Opposite())
				b.prevDirs[next] = d.Opposite()
				b.oxygenPos = next
			default:
				panic("unknown move response")
			}
		}
		switch len(options) {
		case 0:
			if len(branchPoints) == 0 {
				break outer
			}
			lastBranch := branchPoints[len(branchPoints)-1]
			branchPoints = branchPoints[:len(branchPoints)-1]
			for pos != lastBranch {
				prevDir := b.prevDirs[pos]
				b.Move(prevDir)
				pos = pos.Move(prevDir)
			}
		case 1:
			b.Move(options[0])
			pos = pos.Move(options[0])
		default:
			branchPoints = append(branchPoints, pos)
			for i := 1; i < len(options); i++ {
				p := pos.Move(options[i])
				if b.chart[p] == Passable {
					delete(b.chart, p)
				}
			}
			b.Move(options[0])
			pos = pos.Move(options[0])
		}
	}
}

func (b *Bot) RunProg(prog string) int {
	exit := make(chan struct{})

	b.computer.SetInput(prog)

	if os.Getenv("SHOW_GRID") == "true" {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
		b.reallocTermBuffer(termbox.Size())

		go b.eventLoop(exit)
		b.updateFrame()
	}

	go b.computer.Calculate()

	b.exposeMap()
	close(exit)

	<-exit
	return b.oxygenDist
}

func (b *Bot) updateFrame() {
	if os.Getenv("SHOW_GRID") != "true" {
		return
	}
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	copy(termbox.CellBuffer(), b.cells)
	for k, v := range b.chart {
		termbox.SetCell(k.X(), k.Y(), graphics[v], termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.SetCell(b.position.X(), b.position.Y(), graphics[Droid], termbox.ColorWhite, termbox.ColorDefault)
	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (b *Bot) TimeToFillArea() int {
	t := 0
	stack := []grid.Coord{b.oxygenPos}

	for {
		nextStack := []grid.Coord{}
		for _, p := range stack {
			b.chart[p] = Oxygen
			for _, d := range []grid.Direction{grid.North, grid.East, grid.South, grid.West} {
				next := p.Move(d)
				if b.chart[next] == Passable {
					nextStack = append(nextStack, next)
				}
			}
		}
		stack = nextStack
		if len(stack) == 0 {
			break
		}
		t++
	}

	return t
}

func (b *Bot) Move(d grid.Direction) StatusCode {
	b.in <- int(d)
	status := StatusCode(<-b.out)
	curDist := b.dists[b.position]
	nextCoord := b.position.Move(d)

	switch status {
	case HitWall:
		b.chart[nextCoord] = Wall
	case Success:
		b.chart[nextCoord] = Passable
		b.position = nextCoord
		b.dists[nextCoord] = minNotZero(curDist+1, b.dists[nextCoord])
	case FoundOxygen:
		b.chart[nextCoord] = Oxygen
		b.position = nextCoord
		b.dists[nextCoord] = minNotZero(curDist+1, b.dists[nextCoord])
		b.oxygenDist = curDist + 1
	}

	b.updateFrame()
	return status
}

func minNotZero(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
