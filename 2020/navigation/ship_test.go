package navigation_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/navigation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ship", func() {
	var (
		data io.Reader
		ship navigation.Ship
	)

	BeforeEach(func() {
		data = strings.NewReader(`
F10
N3
F7
R90
F11
`)

		ship = navigation.NewShip()
		ship.Load(data)
	})

	It("calcs the correct distance", func() {
		ship.Move()
		Expect(ship.ManhattanDistance()).To(Equal(25))
	})

	It("calcs the correct distance the new way", func() {
		ship.MoveNew()
		Expect(ship.ManhattanDistance()).To(Equal(286))
	})
})
