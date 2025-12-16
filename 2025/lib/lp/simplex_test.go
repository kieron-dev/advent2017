package lp

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructAuxProgram(t *testing.T) {
	constraints := [][]float64{
		{2, -1},
		{1, -5},
	}
	ineqValues := []float64{2, -4}
	objectiveCoeffs := []float64{2, -1}

	l := NewLinearProgram(constraints, ineqValues, objectiveCoeffs)
	aux := l.constructAuxiliaryProgram()

	assert.Equal(t, []float64{-1, 0, 0, 0, 0}, aux.ObjectiveCoeffs)
	assert.Equal(t, 0.0, aux.ObjectiveVal)
	assert.Equal(t, []int{3, 4}, aux.BasicIndices)
	assert.Equal(t, []int{0, 1, 2}, aux.NonBasicIndices)
	assert.Equal(t, [][]float64{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{-1, 2, -1, 0, 0},
		{-1, 1, -5, 0, 0},
	}, aux.Constraints)
	assert.Equal(t, []float64{0, 0, 0, 2, -4}, aux.IneqValues)

	// grim. but...
	aux.Pivot(4, 0)
	assert.Equal(t, []float64{0, -1, 5, 0, -1}, aux.ObjectiveCoeffs)
	assert.Equal(t, -4.0, aux.ObjectiveVal)
	assert.ElementsMatch(t, []int{0, 3}, aux.BasicIndices)
	assert.ElementsMatch(t, []int{1, 2, 4}, aux.NonBasicIndices)
	assert.Equal(t, [][]float64{
		{0, -1, 5, 0, -1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 4, 0, -1},
		{0, 0, 0, 0, 0},
	}, aux.Constraints)
	assert.Equal(t, []float64{4, 0, 0, 6, 0}, aux.IneqValues)

	// also grim
	err := aux.optimise()
	assert.NoError(t, err)
	assert.Equal(t, []float64{-1, 0, 0, 0, 0}, aux.ObjectiveCoeffs)
	assert.Equal(t, 0.0, aux.ObjectiveVal)
	assert.ElementsMatch(t, []int{2, 3}, aux.BasicIndices)
	assert.ElementsMatch(t, []int{0, 1, 4}, aux.NonBasicIndices)
	assert.Equal(t, [][]float64{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0.2, -0.2, 0, 0, -0.2},
		{-0.8, 1.8, 0, 0, -0.2},
		{0, 0, 0, 0, 0},
	}, aux.Constraints)
	assert.Equal(t, []float64{0, 0, 0.8, 2.8, 0}, aux.IneqValues)
}

func TestInitializeSimplex(t *testing.T) {
	constraints := [][]float64{
		{2, -1},
		{1, -5},
	}
	ineqValues := []float64{2, -4}
	objectiveCoeffs := []float64{2, -1}

	lp, err := InitializeSimplex(constraints, ineqValues, objectiveCoeffs)
	assert.NoError(t, err)

	assert.Equal(t, []float64{1.8, 0, 0, -0.2}, lp.ObjectiveCoeffs, "objective coefficients")
	assert.Equal(t, -0.8, lp.ObjectiveVal)
	assert.ElementsMatch(t, []int{1, 2}, lp.BasicIndices)
	assert.ElementsMatch(t, []int{0, 3}, lp.NonBasicIndices)
	assert.Equal(t, [][]float64{
		{0, 0, 0, 0},
		{-0.2, 0, 0, -0.2},
		{1.8, 0, 0, -0.2},
		{0, 0, 0, 0},
	}, lp.Constraints)
	assert.Equal(t, []float64{0, 0.8, 2.8, 0}, lp.IneqValues)
}

func TestSimplex(t *testing.T) {
	constraints := [][]float64{
		{2, -1},
		{1, -5},
	}
	ineqValues := []float64{2, -4}
	objectiveCoeffs := []float64{2, -1}
	_, val, err := Simplex(constraints, ineqValues, objectiveCoeffs)
	assert.NoError(t, err)
	assert.InDelta(t, 2.0, val, 0.000001)
}

func TestAocProblem(t *testing.T) {
	constraints := [][]float64{
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, -1, -1},
		{0, 1, 0, 0, 0, 1},
		{0, -1, 0, 0, 0, -1},
		{0, 0, 1, 1, 1, 0},
		{0, 0, -1, -1, -1, 0},
		{1, 1, 0, 1, 0, 0},
		{-1, -1, 0, -1, 0, 0},
	}
	ineqValues := []float64{3, -3, 5, -5, 4, -4, 7, -7}
	objectiveCoeffs := []float64{-1, -1, -1, -1, -1, -1}
	_, val, err := Simplex(constraints, ineqValues, objectiveCoeffs)
	assert.NoError(t, err)
	assert.Equal(t, -10.0, val)
}

func LineToInput(t *testing.T, line string) ([][]float64, []float64, []float64) {
	firstCurlyBr := strings.Index(line, "{")
	lastCurlyBr := strings.LastIndex(line, "}")
	requireds := []int{}
	bits := strings.Split(line[firstCurlyBr+1:lastCurlyBr], ",")
	for _, b := range bits {
		n, err := strconv.Atoi(b)
		assert.NoError(t, err)
		requireds = append(requireds, n)
	}

	firstRoundBr := strings.Index(line, "(")
	lastRoundBr := strings.LastIndex(line, ")")
	fields := strings.FieldsFunc(line[firstRoundBr:lastRoundBr], func(r rune) bool {
		if r == '(' || r == ')' || r == ' ' {
			return true
		}
		return false
	})
	schematics := [][]int{}
	for _, field := range fields {
		bits := strings.Split(field, ",")
		bitInts := []int{}
		for _, b := range bits {
			n, err := strconv.Atoi(b)
			assert.NoError(t, err)
			bitInts = append(bitInts, n)
		}
		schematics = append(schematics, bitInts)
	}

	constraints := make([][]float64, 2*len(requireds))
	ineqValues := make([]float64, 2*len(requireds))
	objectiveCoeffs := make([]float64, len(schematics))
	for i := range objectiveCoeffs {
		objectiveCoeffs[i] = -1.0
	}
	for i := range constraints {
		constraints[i] = make([]float64, len(schematics))
	}

	for r, v := range requireds {
		for i, s := range schematics {
			if slices.Contains(s, r) {
				constraints[r*2][i] = 1
				constraints[r*2+1][i] = -1
			}
		}
		ineqValues[r*2] = float64(v)
		ineqValues[r*2+1] = -float64(v)
	}
	return constraints, ineqValues, objectiveCoeffs
}

func TestAocProblem2(t *testing.T) {
	line := "[#..##.#..] (1,2,4,6,8) (0,1,6) (0,1,2,5,6,8) (0,1,2,4,6,7,8) (0,2,4,6,7,8) (0,1,2,3,4,5) (3,5,7) (0,3,4,6) {56,47,33,28,34,27,55,4,24}"
	constraints, ineqValues, objectiveCoeffs := LineToInput(t, line)

	s, val, err := Simplex(constraints, ineqValues, objectiveCoeffs)
	assert.NoError(t, err)
	assert.Equal(t, []float64{8, 15, 15, 0, 1, 9, 3, 16}, s)
	assert.Equal(t, -67.0, val)
}

func TestAocProblem3(t *testing.T) {
	line := "[###.] (0,2,3) (0) (1,2) (0,1,3) (0,2) {18,16,28,7}"
	constraints, ineqValues, objectiveCoeffs := LineToInput(t, line)

	s, val, err := Simplex(constraints, ineqValues, objectiveCoeffs)
	assert.NoError(t, err)
	assert.Equal(t, []float64{4, 0, 13, 3, 11}, s)
	assert.Equal(t, -31.0, val)
}
