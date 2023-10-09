package two022_test

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("21", func() {
	example := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

	type statement struct {
		id       string
		val      int
		operandA string
		operandB string
		operator string
	}

	assignRE := regexp.MustCompile(`(\w{4}): (\d+)$`)
	newAssignment := func(s string) *statement {
		matches := assignRE.FindStringSubmatch(s)
		if matches == nil {
			return nil
		}

		n, err := strconv.Atoi(matches[2])
		Expect(err).NotTo(HaveOccurred())

		return &statement{
			id:  matches[1],
			val: n,
		}
	}
	operationRE := regexp.MustCompile(`(\w{4}): (\w{4}) (\S) (\w{4})$`)
	newOperation := func(s string) *statement {
		matches := operationRE.FindStringSubmatch(s)
		if matches == nil {
			return nil
		}

		return &statement{
			id:       matches[1],
			operandA: matches[2],
			operandB: matches[4],
			operator: matches[3],
		}
	}

	load := func(in io.Reader) map[string]*statement {
		ret := map[string]*statement{}

		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			assignment := newAssignment(line)
			if assignment != nil {
				ret[assignment.id] = assignment
				continue
			}

			operation := newOperation(line)
			if operation != nil {
				ret[operation.id] = operation
				continue
			}

			Fail("didn't expect to be here: " + line)
		}

		return ret
	}

	var sort func(map[string]*statement, string, func(*statement))
	sort = func(statements map[string]*statement, start string, cb func(*statement)) {
		el := statements[start]
		if el.operator != "" {
			sort(statements, el.operandA, cb)
			sort(statements, el.operandB, cb)
		}
		cb(el)
	}

	var inSubtree func(map[string]*statement, string, string) bool
	inSubtree = func(statements map[string]*statement, root, el string) bool {
		if root == el {
			return true
		}

		head := statements[root]
		if head.operator == "" {
			return false
		}

		return inSubtree(statements, head.operandA, el) ||
			inSubtree(statements, head.operandB, el)
	}

	calc := func(ordered []*statement, label string) int {
		vals := map[string]int{}
		for _, s := range ordered {
			switch s.operator {
			case "":
				vals[s.id] = s.val
			case "+":
				vals[s.id] = vals[s.operandA] + vals[s.operandB]
			case "-":
				vals[s.id] = vals[s.operandA] - vals[s.operandB]
			case "*":
				vals[s.id] = vals[s.operandA] * vals[s.operandB]
			case "/":
				vals[s.id] = vals[s.operandA] / vals[s.operandB]
			}
		}
		return vals[label]
	}

	var solve func(map[string]*statement, string, string, int) int
	solve = func(statements map[string]*statement, equality, variable string, rval int) int {
		root := statements[equality]

		if root.id == variable {
			return rval
		}

		var val int
		var stmt *statement

		ordered := []*statement{}
		if inSubtree(statements, root.operandA, variable) {
			sort(statements, root.operandB, func(s *statement) {
				ordered = append(ordered, s)
			})
			val = calc(ordered, root.operandB)
			stmt = statements[root.operandA]
			switch root.operator {
			case "+":
				val = rval - val
			case "-":
				val = rval + val
			case "*":
				val = rval / val
			case "/":
				val = rval * val
			}
		} else {
			sort(statements, root.operandA, func(s *statement) {
				ordered = append(ordered, s)
			})
			val = calc(ordered, root.operandA)
			stmt = statements[root.operandB]
			switch root.operator {
			case "+":
				val = rval - val
			case "-":
				val = val - rval
			case "*":
				val = rval / val
			case "/":
				val = val / rval
			}
		}

		return solve(statements, stmt.id, variable, val)
	}

	It("the example", func() {
		statements := load(strings.NewReader(example))

		ordered := []*statement{}
		sort(statements, "root", func(s *statement) {
			ordered = append(ordered, s)
		})

		Expect(calc(ordered, "root")).To(Equal(152))
	})

	It("part A", func() {
		f, err := os.Open("input21")
		Expect(err).NotTo(HaveOccurred())
		statements := load(f)
		f.Close()

		ordered := []*statement{}
		sort(statements, "root", func(s *statement) {
			ordered = append(ordered, s)
		})

		Expect(calc(ordered, "root")).To(Equal(170237589447588))
	})

	It("the example part B", func() {
		statements := load(strings.NewReader(example))
		statements["root"].operator = "-"

		res := solve(statements, "root", "humn", 0)
		Expect(res).To(Equal(301))
	})

	It("part B", func() {
		f, err := os.Open("input21")
		Expect(err).NotTo(HaveOccurred())
		statements := load(f)
		f.Close()

		statements["root"].operator = "-"

		res := solve(statements, "root", "humn", 0)
		Expect(res).To(Equal(3712643961892))
	})
})
