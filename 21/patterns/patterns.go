package patterns

type Art struct {
	pattern []string
}

func New() *Art {
	art := Art{}
	art.pattern = []string{
		".#.",
		"..#",
		"###",
	}
	return &art
}

func (a *Art) Size() int {
	return len(a.pattern)
}

func (a *Art) Pattern() []string {
	return a.pattern
}

func (a *Art) GetKeys(x, y int) []string {
	square := a.GetSquare(x, y)
	out := []string{}
	for i := 0; i < 2; i++ {
		if i == 1 {
			FlipSquare(square)
		}
		for j := 0; j < 4; j++ {
			out = append(out, SquareToString(square))
			RotateSquare(square)
		}
	}
	return out
}

func (a *Art) GetSquare(x, y int) [][]byte {
	l := 3
	if len(a.pattern)%2 == 0 {
		l = 2
	}

	sq := make([][]byte, l)
	for i := 0; i < l; i++ {
		sq[i] = make([]byte, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sq[i][j] = a.pattern[x*l+i][y*l+j]
		}
	}
	return sq
}

func RotateSquare(sq [][]byte) {
	l := len(sq) - 1
	for i := 0; i < l; i++ {
		saved := sq[0][i]
		sq[0][i] = sq[l-i][0]
		sq[l-i][0] = sq[l][l-i]
		sq[l][l-i] = sq[i][l]
		sq[i][l] = saved
	}
}

func FlipSquare(sq [][]byte) {
	l := len(sq)
	for i := 0; i < l; i++ {
		save := sq[i][0]
		sq[i][0] = sq[i][l-1]
		sq[i][l-1] = save
	}
}

func SquareToString(sq [][]byte) string {
	l := len(sq)
	s := ""
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			s += string(sq[i][j])
		}
		if i < l-1 {
			s += "/"
		}
	}
	return s
}
