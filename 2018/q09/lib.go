package q09

type Game struct {
	NumPlayers    int
	NumMarbles    int
	Scores        []int
	currentMarble *Marble
	step          int
}

type Marble struct {
	number int
	next   *Marble
	prev   *Marble
}

func NewGame(nPlayers, nMarbles int) *Game {
	g := Game{
		NumPlayers: nPlayers,
		NumMarbles: nMarbles,
	}
	g.Scores = make([]int, nPlayers)
	zero := Marble{}
	zero.next = &zero
	zero.prev = &zero
	g.currentMarble = &zero
	return &g
}

func (g *Game) Play() int {
	for g.step < g.NumMarbles {
		g.Step()
	}
	maxScore := 0
	for i := 0; i < g.NumPlayers; i++ {
		if g.Scores[i] > maxScore {
			maxScore = g.Scores[i]
		}
	}
	return maxScore
}

func (g *Game) Step() {
	g.step++
	if g.step%23 == 0 {
		g.GivePoints()
	} else {
		m := Marble{
			number: g.step,
		}
		g.InsertMarble(&m)
	}
}

func (g *Game) GivePoints() {
	player := g.step % g.NumPlayers
	if player == 0 {
		player += g.NumPlayers
	}
	g.Scores[player-1] += g.step
	m := g.currentMarble
	for i := 0; i < 7; i++ {
		m = m.prev
	}
	m.prev.next = m.next
	m.next.prev = m.prev
	g.Scores[player-1] += m.number
	g.currentMarble = m.next
}

func (g *Game) InsertMarble(m *Marble) {
	right := g.currentMarble.next
	rightRight := right.next
	m.prev = right
	m.next = rightRight
	right.next = m
	rightRight.prev = m
	g.currentMarble = m
}
