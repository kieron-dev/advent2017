package two023_test

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type rockmap [][]byte

func (r rockmap) print() {
	fmt.Println()
	for _, r := range r {
		fmt.Println(string(r))
	}
	fmt.Println()
}

func (r rockmap) hash() string {
	h := sha256.New()
	for _, row := range r {
		h.Write(row)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (r rockmap) oCount() int {
	s := 0
	for _, row := range r {
		for _, b := range row {
			if b == 'O' {
				s++
			}
		}
	}
	return s
}

func (r rockmap) load() int {
	s := 0
	for i, row := range r {
		for _, b := range row {
			if b == 'O' {
				s += len(r) - i
			}
		}
	}
	return s
}

func (r rockmap) tiltNorth() {
	for row := 1; row < len(r); row++ {
		for col := 0; col < len(r[0]); col++ {
			char := r[row][col]
			if char != 'O' {
				continue
			}
			for k := row - 1; k >= -1; k-- {
				if k == -1 || r[k][col] != '.' {
					if r[k+1][col] == '.' {
						r[k+1][col] = 'O'
						r[row][col] = '.'
					}
					break
				}
			}
		}
	}
}

func (r rockmap) tiltSouth() {
	for row := len(r) - 2; row >= 0; row-- {
		for col := 0; col < len(r[0]); col++ {
			char := r[row][col]
			if char != 'O' {
				continue
			}
			for k := row + 1; k <= len(r); k++ {
				if k == len(r) || r[k][col] != '.' {
					if r[k-1][col] == '.' {
						r[k-1][col] = 'O'
						r[row][col] = '.'
					}
					break
				}
			}
		}
	}
}

func (r rockmap) tiltWest() {
	for col := 1; col < len(r[0]); col++ {
		for row := 0; row < len(r); row++ {
			char := r[row][col]
			if char != 'O' {
				continue
			}
			for k := col - 1; k >= -1; k-- {
				if k == -1 || r[row][k] != '.' {
					if r[row][k+1] == '.' {
						r[row][k+1] = 'O'
						r[row][col] = '.'
					}
					break
				}
			}
		}
	}
}

func (r rockmap) tiltEast() {
	for col := len(r[0]) - 1; col >= 0; col-- {
		for row := 0; row < len(r); row++ {
			char := r[row][col]
			if char != 'O' {
				continue
			}
			for k := col + 1; k <= len(r[0]); k++ {
				if k == len(r[0]) || r[row][k] != '.' {
					if r[row][k-1] == '.' {
						r[row][k-1] = 'O'
						r[row][col] = '.'
					}
					break
				}
			}
		}
	}
}

func (r rockmap) cycle() {
	r.tiltNorth()
	r.tiltWest()
	r.tiltSouth()
	r.tiltEast()
}

func loadRockMap(filename string) rockmap {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	r := rockmap{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			r = append(r, []byte(line))
		}
	}
	return r
}

var _ = Describe("14", func() {
	It("does part A", func() {
		r := loadRockMap("input14")
		r.tiltNorth()
		Expect(r.load()).To(Equal(106997))
	})

	It("does part B", func() {
		r := loadRockMap("input14")
		hashes := map[string]int{}
		i := 0
		var start, interval int
		for {
			r.cycle()
			h := r.hash()
			prev, ok := hashes[h]
			if ok {
				start = prev
				interval = i - prev
				break
			}
			hashes[h] = i
			i++
		}

		extra := (1_000_000_000 - start) % interval
		for i := 0; i < extra-1; i++ {
			r.cycle()
		}

		Expect(r.load()).To(Equal(99641))
	})
})
