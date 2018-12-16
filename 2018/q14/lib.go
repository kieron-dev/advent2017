package q14

import (
	"bytes"
	"fmt"
	"strings"
)

type Recipes struct {
	Start      *Recipe
	ElfARecipe *Recipe
	ElfBRecipe *Recipe
	End        *Recipe
	Length     int
}

type Recipe struct {
	Score int
	Left  *Recipe
	Right *Recipe
}

func NewRecipes() *Recipes {
	three := Recipe{
		Score: 3,
	}
	seven := Recipe{
		Score: 7,
	}
	three.Right = &seven
	three.Left = &seven
	seven.Right = &three
	seven.Left = &three
	r := Recipes{
		Start:      &three,
		ElfARecipe: &three,
		ElfBRecipe: &seven,
		End:        &seven,
		Length:     2,
	}
	return &r
}

func (r *Recipes) ScoresBefore(sequence string) int {
	for {
		for i := 0; i < 10000000; i++ {
			r.Step()
		}
		fmt.Printf("r.Length = %+v\n", r.Length)
		scores := r.String()
		pos := strings.Index(scores, sequence)
		if pos > -1 {
			return pos
		}
	}
}

func (r *Recipes) String() string {
	buf := bytes.Buffer{}
	node := r.Start
	for {
		buf.WriteString(fmt.Sprintf("%d", node.Score))
		node = node.Right
		if node == r.Start {
			break
		}
	}
	return buf.String()
}

func (r *Recipes) ScoresAfter(n int) string {
	for r.Length < n+10 {
		r.Step()
	}
	node := r.End
	if r.Length == n+11 {
		node = node.Left
	}
	res := ""
	for i := 0; i < 10; i++ {
		res = fmt.Sprintf("%d", node.Score) + res
		node = node.Left
	}
	return res
}

func (r *Recipes) Step() {
	sum := r.ElfARecipe.Score + r.ElfBRecipe.Score
	if sum > 9 {
		NewRecipe := Recipe{Score: 1}
		r.Append(&NewRecipe)
	}
	Recipe := Recipe{Score: sum % 10}
	r.Append(&Recipe)
	scoreA := r.ElfARecipe.Score
	for i := 0; i < scoreA+1; i++ {
		r.ElfARecipe = r.ElfARecipe.Right
	}
	scoreB := r.ElfBRecipe.Score
	for i := 0; i < scoreB+1; i++ {
		r.ElfBRecipe = r.ElfBRecipe.Right
	}
}

func (r *Recipes) Append(newRecipe *Recipe) {
	r.End.Right = newRecipe
	r.Start.Left = newRecipe
	newRecipe.Left = r.End
	newRecipe.Right = r.Start
	r.End = newRecipe
	r.Length++
}
