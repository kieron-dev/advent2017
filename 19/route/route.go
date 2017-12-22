package route

var (
	Down  = Point{r: 1, c: 0}
	Up    = Point{r: -1, c: 0}
	Left  = Point{r: 0, c: -1}
	Right = Point{r: 0, c: 1}
)

type Map struct {
	rows    []string
	curPos  Point
	curDir  Point
	letters string
	steps   int
}

type Point struct {
	r int
	c int
}

func (p Point) Add(d Point) Point {
	return Point{r: p.r + d.r, c: p.c + d.c}
}

func NewPoint(r, c int) Point {
	return Point{r: r, c: c}
}

func New(rows []string) *Map {
	m := Map{rows: rows}
	for i, c := range rows[0] {
		if c != ' ' {
			m.curPos = Point{r: 0, c: i}
			break
		}
	}
	m.curDir = Down
	m.steps = 1
	return &m
}

func (m *Map) CurPosition() Point {
	return m.curPos
}

func (m *Map) CurDirection() Point {
	return m.curDir
}

func (m *Map) CharAt(p Point) byte {
	if p.r < 0 || p.r >= len(m.rows) || p.c < 0 || p.c >= len(m.rows[p.r]) {
		return ' '
	}
	return m.rows[p.r][p.c]
}

func (m *Map) GetLeftRightRelativeDirections() []Point {
	if m.curDir == Up || m.curDir == Down {
		return []Point{Left, Right}
	}
	return []Point{Up, Down}
}

func (m *Map) Step() bool {
	next := m.curPos.Add(m.curDir)
	nextChar := m.CharAt(next)
	if nextChar != ' ' {
		m.curPos = next
		m.recordChar(nextChar)
		m.steps++
		return true
	}
	for _, d := range m.GetLeftRightRelativeDirections() {
		try := m.curPos.Add(d)
		nextChar = m.CharAt(try)
		if nextChar != ' ' {
			m.curPos = try
			m.curDir = d
			m.recordChar(nextChar)
			m.steps++
			return true
		}
	}
	return false
}

func (m *Map) recordChar(c byte) {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		m.letters += string(c)
	}
}

func (m *Map) Walk() {
	for {
		if !m.Step() {
			break
		}
	}
}

func (m *Map) GetLetters() string {
	return m.letters
}

func (m *Map) GetStepCount() int {
	return m.steps
}
