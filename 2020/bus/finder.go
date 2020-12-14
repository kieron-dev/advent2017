package bus

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/maths"
)

type Finder struct {
	start int
	buses map[int]int
}

func NewFinder() Finder {
	return Finder{
		buses: map[int]int{},
	}
}

func (f *Finder) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	if !scanner.Scan() {
		log.Fatalf("expected to read a line")
	}
	line1 := scanner.Text()
	line1 = strings.TrimSpace(line1)

	var err error
	f.start, err = strconv.Atoi(line1)
	if err != nil {
		log.Fatalf("failed to parse start time %q: %v", line1, err)
	}

	if !scanner.Scan() {
		log.Fatalf("expected to read a line")
	}

	line2 := scanner.Text()
	line2 = strings.TrimSpace(line2)

	for i, b := range strings.Split(line2, ",") {
		if b == "x" {
			continue
		}

		n, err := strconv.Atoi(b)
		if err != nil {
			log.Fatalf("getting bus no failed: %v", err)
		}
		f.buses[n] = (n - i) % n
	}
}

func (f Finder) Find() (busNumber, waitMins int) {
	minWait := f.start
	busNo := -1

	for b := range f.buses {
		if b == 0 {
			continue
		}

		wait := b - (f.start % b)
		if wait < minWait {
			minWait = wait
			busNo = b
		}
	}

	return busNo, minWait
}

func (f Finder) SpecialTimestamp() int {
	workMap := make(map[*big.Int]*big.Int, len(f.buses))

	for k, v := range f.buses {
		workMap[big.NewInt(int64(k))] = big.NewInt(int64(v))
	}

	for len(workMap) > 1 {
		var i int
		var busNos, remainders [2]*big.Int

		for busNo, remainder := range workMap {
			if i == 2 {
				break
			}

			busNos[i] = busNo
			remainders[i] = remainder

			i++
		}

		d, s, t := maths.ExtEuclid(busNos[0], busNos[1])

		if d.Int64() != 1 {
			log.Fatalf("strategy failed! nums not coprime: %v", busNos)
		}

		newRemainder0 := new(big.Int).Mul(new(big.Int).Mul(remainders[0], t), busNos[1])
		newRemainder1 := new(big.Int).Mul(new(big.Int).Mul(remainders[1], s), busNos[0])
		newRemainder2 := new(big.Int).Add(newRemainder0, newRemainder1)

		busProd := new(big.Int).Mul(busNos[0], busNos[1])
		newRemainder := new(big.Int).Mod(newRemainder2, busProd)

		if newRemainder.Cmp(big.NewInt(0)) < 0 {
			newRemainder = new(big.Int).Add(newRemainder, busProd)
		}

		delete(workMap, busNos[0])
		delete(workMap, busNos[1])
		workMap[busProd] = newRemainder
	}

	fmt.Printf("workMap = %+v\n", workMap)
	for _, v := range workMap {
		return int(v.Int64())
	}

	log.Fatal("eek")
	return 0
}
