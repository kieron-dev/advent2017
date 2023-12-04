package two023_test

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("03", func() {
	type num struct {
		n        int
		row      int
		colStart int
		// not inclusive
		colEnd int
	}

	getNums := func(rows [][]byte, r int) []num {
		nums := []num{}

		inNum := false
		var n int
		var start, end int

		for i, b := range rows[r] {
			switch {
			case '0' <= b && b <= '9':
				if !inNum {
					start = i
					inNum = true
				}
				n = 10*n + int(b-'0')
			default:
				if inNum {
					end = i - 1
					inNum = false
					nums = append(nums, num{n: n, row: r, colStart: start, colEnd: end})
				}
				n = 0
			}
		}
		if inNum {
			nums = append(nums, num{n: n, row: r, colStart: start, colEnd: len(rows[r]) - 1})
		}

		return nums
	}

	isSymbol := func(rows [][]byte, r, c int) bool {
		if r < 0 || r >= len(rows) || c < 0 || c >= len(rows[0]) {
			return false
		}
		b := rows[r][c]
		return b != '.' && (b < '0' || b > '9')
	}

	isPart := func(n num, rows [][]byte) bool {
		for row := n.row - 1; row <= n.row+1; row++ {
			for col := n.colStart - 1; col <= n.colEnd+1; col++ {
				if isSymbol(rows, row, col) {
					return true
				}
			}
		}

		return false
	}

	getStars := func(rows [][]byte, r int) []int {
		res := []int{}
		for i := range rows[r] {
			if rows[r][i] == '*' {
				res = append(res, i)
			}
		}

		return res
	}

	isAdjacent := func(n num, r, c int) bool {
		if n.row == r {
			return n.colStart-1 == c || n.colEnd+1 == c
		}

		return n.colStart-1 <= c && c <= n.colEnd+1
	}

	gearVal := func(rows [][]byte, r, c int) int {
		nums := []num{}
		if r > 0 {
			for _, num := range getNums(rows, r-1) {
				if isAdjacent(num, r, c) {
					nums = append(nums, num)
				}
			}
		}
		for _, num := range getNums(rows, r) {
			if isAdjacent(num, r, c) {
				nums = append(nums, num)
			}
		}
		if r < len(rows)-1 {
			for _, num := range getNums(rows, r+1) {
				if isAdjacent(num, r, c) {
					nums = append(nums, num)
				}
			}
		}

		if len(nums) != 2 {
			return 0
		}

		return nums[0].n * nums[1].n
	}

	It("does part A", func() {
		bs, err := os.ReadFile("input03")
		Expect(err).NotTo(HaveOccurred())

		rows := bytes.Fields(bs)

		sum := 0
		for i := range rows {
			for _, num := range getNums(rows, i) {
				if isPart(num, rows) {
					sum += num.n
				}
			}
		}

		Expect(sum).To(Equal(557705))
	})

	It("does part B", func() {
		bs, err := os.ReadFile("input03")
		Expect(err).NotTo(HaveOccurred())

		rows := bytes.Fields(bs)

		sum := 0
		for i := range rows {
			for _, col := range getStars(rows, i) {
				sum += gearVal(rows, i, col)
			}
		}

		Expect(sum).To(Equal(84266818))
	})
})
