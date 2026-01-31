// Package lp is a Linear Programming package.
// It implements the Simplex algorithm.
package lp

import (
	"errors"
	"slices"
)

// Use small or -small when comparing values with zero
const small = 1e-8

type LinearProgram struct {
	// N
	NonBasicIndices []int
	// B
	BasicIndices []int
	// A
	Constraints [][]float64
	// b
	IneqValues []float64
	// c
	ObjectiveCoeffs []float64
	// v
	ObjectiveVal float64
}

func NewLinearProgram(constraints [][]float64, ineqValues []float64, objectiveCoeffs []float64) *LinearProgram {
	l := new(LinearProgram)

	// n is the number of non-basic variables, i.e. x_0, ..., x_{n-1}
	n := len(constraints[0])
	// m is the number of basic variables, i.e. slack variables x_n, ..., x_{m-1}
	m := len(constraints)

	// constraints is a n (cols) x m (rows) matrix: n vars, m constraints.
	// We insert n rows to hold inequalities for the original non-basic vars at the beginning
	l.Constraints = make([][]float64, n+m)
	for i := range l.Constraints {
		l.Constraints[i] = make([]float64, n+m)
	}
	for i := n; i < len(l.Constraints); i++ {
		copy(l.Constraints[i], constraints[i-n])
	}

	// inequality values start off for the slack variables n, n+1, ...n n+m-1.
	l.IneqValues = make([]float64, n, n+m)
	l.IneqValues = append(l.IneqValues, ineqValues...)

	// objective coefficients start off for the non-basic vars 0, ..., n-1
	l.ObjectiveCoeffs = make([]float64, n+m)
	copy(l.ObjectiveCoeffs, objectiveCoeffs)

	l.NonBasicIndices = make([]int, n)
	for i := range l.NonBasicIndices {
		l.NonBasicIndices[i] = i
	}

	l.BasicIndices = make([]int, m)
	for i := range l.BasicIndices {
		l.BasicIndices[i] = n + i
	}

	return l
}

func InitializeSimplex(constraints [][]float64, ineqValues []float64, objectiveCoeffs []float64) (*LinearProgram, error) {
	if len(constraints) != len(ineqValues) {
		panic("on the streets of london")
	}
	l := NewLinearProgram(constraints, ineqValues, objectiveCoeffs)

	minIneqVal := l.IneqValues[0]
	minIneqIdx := 0
	for i, b := range l.IneqValues {
		if b < minIneqVal {
			minIneqVal = b
			minIneqIdx = i
		}
	}

	if minIneqVal >= -small {
		// initial basic solution is feasible
		return l, nil
	}

	lAux := l.constructAuxiliaryProgram()

	// we've inserted 0 at the front, so add 1 to leaving index
	lAux.Pivot(minIneqIdx+1, 0)

	err := lAux.optimise()
	if err != nil {
		return nil, err
	}

	x0 := lAux.IneqValues[0]
	if slices.Contains(lAux.BasicIndices, 0) {
		x0 = 0.0
	}
	if x0 != 0 {
		return nil, errors.New("infeasible")
	}

	// ensure x_0 non-basic. If basic, perform degenerate pivot to make it non-basic
	if slices.Contains(lAux.BasicIndices, 0) {
		enteringIndex := -1
		for i, v := range lAux.ObjectiveCoeffs {
			if v < -small {
				enteringIndex = i
				break
			}
		}
		if enteringIndex == -1 {
			return nil, errors.New("can't do a degenerate pivot")
		}
		lAux.Pivot(0, enteringIndex)
	}

	for i, c := range lAux.Constraints {
		lAux.Constraints[i] = c[1:]
	}
	lAux.Constraints = lAux.Constraints[1:]

	for i, b := range lAux.BasicIndices {
		lAux.BasicIndices[i] = b - 1
	}
	for i, n := range lAux.NonBasicIndices {
		lAux.NonBasicIndices[i] = n - 1
	}
	idx := slices.Index(lAux.NonBasicIndices, -1)
	lAux.NonBasicIndices = append(lAux.NonBasicIndices[:idx], lAux.NonBasicIndices[idx+1:]...)

	lAux.IneqValues = lAux.IneqValues[1:]

	// restore original objective function, but swap each basic variable with the RHS of its constraint
	lAux.ObjectiveCoeffs = l.ObjectiveCoeffs
	lAux.ObjectiveVal = 0
	for i, coeff := range lAux.ObjectiveCoeffs {
		if !slices.Contains(lAux.BasicIndices, i) {
			continue
		}
		lAux.ObjectiveCoeffs[i] = 0
		lAux.ObjectiveVal += coeff * lAux.IneqValues[i]
		for j, c := range lAux.Constraints[i] {
			lAux.ObjectiveCoeffs[j] += coeff * -c
		}
	}

	return lAux, nil
}

// constraint auxiliary program set to optimise for -x_0
func (lp *LinearProgram) constructAuxiliaryProgram() *LinearProgram {
	n := len(lp.NonBasicIndices)
	m := len(lp.BasicIndices)

	l := new(LinearProgram)

	l.Constraints = make([][]float64, n+m+1)
	l.Constraints[0] = make([]float64, n+m+1)
	// add -x_0 to left of constraints
	for i := range l.Constraints {
		if i <= n {
			l.Constraints[i] = make([]float64, n+m+1)
		} else {
			l.Constraints[i] = append([]float64{-1}, lp.Constraints[i-1]...)
		}
	}

	l.IneqValues = make([]float64, n+m+1)
	for i := n + 1; i <= n+m; i++ {
		l.IneqValues[i] = lp.IneqValues[i-1]
	}

	// objective starts at -x_0
	l.ObjectiveCoeffs = make([]float64, n+m+1)
	l.ObjectiveCoeffs[0] = -1

	l.NonBasicIndices = make([]int, n+1)
	for i := range l.NonBasicIndices {
		l.NonBasicIndices[i] = i
	}

	l.BasicIndices = make([]int, m)
	for i := range len(l.BasicIndices) {
		l.BasicIndices[i] = n + i + 1
	}

	return l
}

func Simplex(constraints [][]float64, ineqValues []float64, objectiveCoeffs []float64) ([]float64, float64, error) {
	l, err := InitializeSimplex(constraints, ineqValues, objectiveCoeffs)
	if err != nil {
		return nil, 0, err
	}
	err = l.optimise()
	if err != nil {
		return nil, 0, err
	}
	ret := make([]float64, len(constraints[0]))
	for i := range ret {
		ret[i] = 0
		if slices.Contains(l.BasicIndices, i) {
			ret[i] = l.IneqValues[i]
		}
	}
	return ret, l.ObjectiveVal, nil
}

func (lp *LinearProgram) optimise() error {
	for {
		e := -1
		for i, c := range lp.ObjectiveCoeffs {
			if c > small {
				e = i
				break
			}
		}
		if e == -1 {
			return nil
		}

		l := -1
		minVal := 1e100
		for _, i := range lp.BasicIndices {
			if lp.Constraints[i][e] > small {
				work := lp.IneqValues[i] / lp.Constraints[i][e]
				if work < minVal {
					minVal = work
					l = i
				}
			}
		}
		if l == -1 {
			return errors.New("unbounded")
		}

		lp.Pivot(l, e)
	}
}

func (lp *LinearProgram) Pivot(leavingIndex int, enteringIndex int) {
	newConstraints := make([][]float64, len(lp.Constraints))
	for i := range newConstraints {
		newConstraints[i] = make([]float64, len(lp.Constraints[i]))
	}

	newIneqValues := make([]float64, len(lp.IneqValues))
	newIneqValues[enteringIndex] = lp.IneqValues[leavingIndex] / lp.Constraints[leavingIndex][enteringIndex]

	for _, j := range lp.NonBasicIndices {
		if j == enteringIndex {
			continue
		}
		newConstraints[enteringIndex][j] = lp.Constraints[leavingIndex][j] / lp.Constraints[leavingIndex][enteringIndex]
	}
	newConstraints[enteringIndex][leavingIndex] = 1 / lp.Constraints[leavingIndex][enteringIndex]

	for _, i := range lp.BasicIndices {
		if i == leavingIndex {
			continue
		}
		newIneqValues[i] = lp.IneqValues[i] - lp.Constraints[i][enteringIndex]*newIneqValues[enteringIndex]
		for _, j := range lp.NonBasicIndices {
			if j == enteringIndex {
				continue
			}
			newConstraints[i][j] = lp.Constraints[i][j] - lp.Constraints[i][enteringIndex]*newConstraints[enteringIndex][j]
		}
		newConstraints[i][leavingIndex] = -lp.Constraints[i][enteringIndex] * newConstraints[enteringIndex][leavingIndex]
	}

	newObjectiveVal := lp.ObjectiveVal + lp.ObjectiveCoeffs[enteringIndex]*newIneqValues[enteringIndex]
	newObjectiveCoeffs := make([]float64, len(lp.ObjectiveCoeffs))
	for _, j := range lp.NonBasicIndices {
		if j == enteringIndex {
			continue
		}
		newObjectiveCoeffs[j] = lp.ObjectiveCoeffs[j] - lp.ObjectiveCoeffs[enteringIndex]*newConstraints[enteringIndex][j]
	}
	newObjectiveCoeffs[leavingIndex] = -lp.ObjectiveCoeffs[enteringIndex] * newConstraints[enteringIndex][leavingIndex]

	newNonBasicIndices := make([]int, len(lp.NonBasicIndices))
	copy(newNonBasicIndices, lp.NonBasicIndices)
	for i, v := range newNonBasicIndices {
		if v == enteringIndex {
			newNonBasicIndices[i] = leavingIndex
		}
	}

	newBasicIndices := make([]int, len(lp.BasicIndices))
	copy(newBasicIndices, lp.BasicIndices)
	for i, v := range newBasicIndices {
		if v == leavingIndex {
			newBasicIndices[i] = enteringIndex
		}
	}

	lp.NonBasicIndices = newNonBasicIndices
	lp.BasicIndices = newBasicIndices
	lp.Constraints = newConstraints
	lp.IneqValues = newIneqValues
	lp.ObjectiveCoeffs = newObjectiveCoeffs
	lp.ObjectiveVal = newObjectiveVal
}
