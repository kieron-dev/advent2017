// Package memory simulates an elf memory game
package memory

import (
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type pair struct {
	latest   int
	previous int
}

type Game struct {
	lastSaid map[int]pair
	words    []int
	pos      int
}

func NewGame() Game {
	return Game{
		lastSaid: map[int]pair{},
	}
}

func (g *Game) Load(data io.Reader) {
	contents, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatalf("read error: %v", err)
	}

	entries := strings.Split(string(contents), ",")

	for _, s := range entries {
		s = strings.TrimSpace(s)
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("expected an int, got %q: %v", s, err)
		}

		g.Add(n)
	}
}

func (g *Game) Add(n int) {
	g.pos++
	g.words = append(g.words, n)
	g.lastSaid[n] = pair{latest: g.pos}
}

func (g *Game) Get(pos int) int {
	for pos > g.pos {
		last := g.words[g.pos-1]
		lastSaid := g.lastSaid[last]
		n := 0
		if lastSaid.previous != 0 {
			n = lastSaid.latest - lastSaid.previous
		}
		g.pos++
		g.words = append(g.words, n)

		nLastSaid := g.lastSaid[n]
		g.lastSaid[n] = pair{latest: g.pos, previous: nLastSaid.latest}
	}

	return g.words[pos-1]
}
