package q15

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

type Fight struct {
	Grid   [][]rune
	Health map[Coord]int
	ElfHit int
}

type Coord struct {
	Row int
	Col int
}

func (f *Fight) AdjacentCells(coord Coord) []Coord {
	cells := []Coord{}
	for _, delta := range []Coord{
		{Row: -1, Col: 0},
		{Row: 1, Col: 0},
		{Row: 0, Col: -1},
		{Row: 0, Col: 1},
	} {
		cell := Coord{Row: coord.Row + delta.Row, Col: coord.Col + delta.Col}
		if cell.Row < 0 || cell.Row >= len(f.Grid) || cell.Col < 0 || cell.Col >= len(f.Grid[0]) {
			continue
		}
		cells = append(cells, cell)
	}
	return cells
}

func (f *Fight) SetElfHit(hit int) {
	f.ElfHit = hit
}

func NewFight(in io.Reader) *Fight {
	f := Fight{}
	f.Grid = [][]rune{}
	br := bufio.NewReader(in)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Trim(line, "\n")
		f.Grid = append(f.Grid, []rune(line))
	}
	f.Health = map[Coord]int{}
	for _, c := range f.GetActorCoords() {
		f.Health[c] = 200
	}
	return &f
}

func (f *Fight) At(coord Coord) rune {
	return f.Grid[coord.Row][coord.Col]
}

func (f *Fight) GetActorCoords() []Coord {
	coords := []Coord{}
	for r := 0; r < len(f.Grid); r++ {
		for c := 0; c < len(f.Grid[0]); c++ {
			coord := Coord{Row: r, Col: c}
			cell := f.At(coord)
			if cell == 'E' || cell == 'G' {
				coords = append(coords, coord)
			}
		}
	}
	return coords
}

func (f *Fight) GetEnemy(from Coord) rune {
	var enemy rune
	switch f.At(from) {
	case 'E':
		enemy = 'G'
	case 'G':
		enemy = 'E'
	default:
		panic("fuck")
	}
	return enemy
}

func (f *Fight) GetAttackSquares(from Coord) []Coord {
	coords := map[Coord]bool{}
	enemy := f.GetEnemy(from)

	for r := 0; r < len(f.Grid); r++ {
		for c := 0; c < len(f.Grid[0]); c++ {
			coord := Coord{Row: r, Col: c}
			cell := f.At(coord)
			if cell == enemy {
				for _, adjacent := range f.AdjacentCells(coord) {
					if f.At(adjacent) == '.' {
						coords[adjacent] = true
					}
				}
			}
		}
	}

	coordList := []Coord{}
	for coord, _ := range coords {
		coordList = append(coordList, coord)
	}
	sort.Slice(coordList, func(i, j int) bool {
		a := coordList[i]
		b := coordList[j]
		if a.Row == b.Row {
			return a.Col < b.Col
		}
		return a.Row < b.Row
	})

	return coordList
}

func IsIn(coords []Coord, coord Coord) bool {
	for _, c := range coords {
		if coord == c {
			return true
		}
	}
	return false
}

func (f *Fight) NearestAttack(from Coord) Coord {
	attackSquares := f.GetAttackSquares(from)
	distances := map[Coord]int{}

	visited := map[Coord]bool{}
	distances[from] = 0
	queue := []Coord{from}
	minDistance := 99999999

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		if visited[cell] {
			continue
		}

		if IsIn(attackSquares, cell) {
			if distances[cell] > minDistance {
				break
			}
			minDistance = distances[cell]
		}

		for _, adj := range f.AdjacentCells(cell) {
			if visited[adj] || f.At(adj) != '.' {
				continue
			}
			distances[adj] = distances[cell] + 1
			queue = append(queue, adj)
		}
		visited[cell] = true
	}

	sort.Slice(attackSquares, func(i, j int) bool {
		if distances[attackSquares[i]] == distances[attackSquares[j]] {
			if attackSquares[i].Row == attackSquares[j].Row {
				return attackSquares[i].Col < attackSquares[j].Col
			}
			return attackSquares[i].Row < attackSquares[j].Row
		}
		return distances[attackSquares[i]] < distances[attackSquares[j]]
	})

	for _, cell := range attackSquares {
		if _, ok := distances[cell]; ok {
			return cell
		}
	}
	return Coord{}
}

func (f *Fight) NextSquare(attacker, target Coord) Coord {
	adjacentCells := f.AdjacentCells(attacker)
	distances := map[Coord]int{}
	for _, adj := range adjacentCells {
		if f.At(adj) != '.' {
			continue
		}
		distances[adj] = f.Distance(adj, target)
	}
	sort.Slice(adjacentCells, func(i, j int) bool {
		if distances[adjacentCells[i]] == distances[adjacentCells[j]] {
			if adjacentCells[i].Row == adjacentCells[j].Row {
				return adjacentCells[i].Col < adjacentCells[j].Col
			}
			return adjacentCells[i].Row < adjacentCells[j].Row
		}
		return distances[adjacentCells[i]] < distances[adjacentCells[j]]
	})
	for _, cell := range adjacentCells {
		if _, ok := distances[cell]; ok {
			return cell
		}
	}
	return Coord{}
}

func (f *Fight) Distance(from, to Coord) int {
	distances := map[Coord]int{}

	visited := map[Coord]bool{}
	distances[from] = 0
	queue := []Coord{from}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		if visited[cell] {
			continue
		}
		if cell == to {
			return distances[cell]
		}
		for _, adj := range f.AdjacentCells(cell) {
			if visited[adj] || f.At(adj) != '.' {
				continue
			}
			distances[adj] = distances[cell] + 1
			queue = append(queue, adj)
		}
		visited[cell] = true
	}
	return 99999999
}

func (f *Fight) GetVictim(c Coord) Coord {
	enemies := []Coord{}
	enemy := f.GetEnemy(c)
	for _, adj := range f.AdjacentCells(c) {
		if f.At(adj) == enemy {
			enemies = append(enemies, adj)
		}
	}

	sort.Slice(enemies, func(i, j int) bool {
		a := enemies[i]
		b := enemies[j]
		if f.Health[a] == f.Health[b] {
			if a.Row == b.Row {
				return a.Col < b.Col
			}
			return a.Row < b.Row
		}
		return f.Health[a] < f.Health[b]
	})
	return enemies[0]
}

func (f *Fight) IsAttacking(c Coord) bool {
	enemy := f.GetEnemy(c)
	for _, adj := range f.AdjacentCells(c) {
		if f.At(adj) == enemy {
			return true
		}
	}
	return false
}

func (f *Fight) EnemyEliminated(c Coord) bool {
	enemy := f.GetEnemy(c)
	for r := 0; r < len(f.Grid); r++ {
		for c := 0; c < len(f.Grid[0]); c++ {
			if f.Grid[r][c] == enemy {
				return false
			}
		}
	}
	return true
}

func RunWithNoElfDeath(in io.Reader) int {
	bs, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(bs)
	f := NewFight(buf)
	elfCount := f.ElfCount()
	elfPower := 4
	lastFail := 4
	lastSuccess := 10000000
	lastSuccessRes := 0

	for {
		fmt.Printf("elfPower = %+v\n", elfPower)
		buf := bytes.NewBuffer(bs)
		f := NewFight(buf)
		f.SetElfHit(elfPower)
		res := f.Run()
		fmt.Printf("res = %+v\n", res)
		newElfCount := f.ElfCount()
		if newElfCount < elfCount {
			lastFail = elfPower
			add := (lastSuccess - elfPower) / 2
			if elfPower < add {
				add = elfPower
			}
			elfPower += add
		} else {
			lastSuccess = elfPower
			lastSuccessRes = res
			elfPower -= (elfPower - lastFail) / 2
		}
		if lastSuccess == lastFail+1 {
			return lastSuccessRes
		}
	}
}

func (f *Fight) Run() int {
	i := 0
	for {
		if !f.Step() {
			break
		}
		f.Print()
		fmt.Println()
		time.Sleep(time.Millisecond * 80)
		i++
	}
	f.Print()
	sum := 0
	for _, h := range f.Health {
		if h > 0 {
			sum += h
		}
	}
	return i * sum
}

func (f *Fight) GetHitPower(c Coord) int {
	switch f.At(c) {
	case 'G':
		return 3
	case 'E':
		return f.ElfHit
	default:
		return 0
	}
}

func (f *Fight) ElfCount() int {
	count := 0
	for r := 0; r < len(f.Grid); r++ {
		for c := 0; c < len(f.Grid[0]); c++ {
			if f.At(Coord{Row: r, Col: c}) == 'E' {
				count++
			}
		}
	}
	return count
}

func (f *Fight) Step() bool {
	for _, actor := range f.GetActorCoords() {
		if f.At(actor) == '.' {
			// killed in a previous step
			continue
		}
		hitPower := f.GetHitPower(actor)
		if f.IsAttacking(actor) {
			victim := f.GetVictim(actor)
			f.Attack(victim, hitPower)
			continue
		}
		t := f.At(actor)
		if f.EnemyEliminated(actor) {
			return false
		}
		target := f.NearestAttack(actor)
		blank := Coord{}
		if target == blank {
			continue
		}
		nextSquare := f.NextSquare(actor, target)
		f.Set(actor, '.')
		f.Set(nextSquare, t)
		f.Health[nextSquare] = f.Health[actor]
		delete(f.Health, actor)

		if f.IsAttacking(nextSquare) {
			victim := f.GetVictim(nextSquare)
			f.Attack(victim, hitPower)
			continue
		}
	}
	return true
}

func (f *Fight) Attack(victim Coord, hit int) {
	f.Health[victim] -= hit
	if f.Health[victim] <= 0 {
		f.Set(victim, '.')
	}
}

func (f *Fight) Set(coord Coord, val rune) {
	f.Grid[coord.Row][coord.Col] = val
}

func (f *Fight) Print() {
	for r := 0; r < len(f.Grid); r++ {
		for c := 0; c < len(f.Grid[0]); c++ {
			fmt.Printf("%s", string(f.Grid[r][c]))
		}
		fmt.Println()
	}
}
