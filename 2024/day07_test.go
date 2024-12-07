package twenty24

import (
	"bufio"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test07A(t *testing.T) {
	input, err := os.Open("input07")
	assert.NoError(t, err)
	defer input.Close()

	sumA := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ":", "", 1)
		nums := asToIs(t, line)

		target := nums[0]

		for p := range int(math.Pow(2, float64(len(nums)-2))) {
			// for p := range 1 << (len(nums) - 2) {
			s := nums[1]
			for i, n := range nums[2:] {
				if p&(1<<i) == 0 {
					s += n
				} else {
					s *= n
				}
				if s > target {
					break
				}
			}
			if s == target {
				sumA += target
				break
			}
		}
	}
	assert.Equal(t, 850435817339, sumA)
}

func Test07B(t *testing.T) {
	input, err := os.Open("input07")
	assert.NoError(t, err)
	defer input.Close()

	sumB := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ":", "", 1)
		nums := asToIs(t, line)

		target := nums[0]

		for p := range int(math.Pow(3, float64(len(nums)-2))) {
			s := nums[1]
			b3 := p
			for _, n := range nums[2:] {
				switch b3 % 3 {
				case 0:
					s += n
				case 1:
					s *= n
				case 2:
					c := n
					for c > 0 {
						s *= 10
						c /= 10
					}
					s += n
				}
				if s > target {
					break
				}
				b3 /= 3
			}
			if s == target {
				sumB += target
				break
			}
		}
	}
	assert.Equal(t, 104824810233437, sumB)
}
