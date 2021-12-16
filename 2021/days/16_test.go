package days_test

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type bitComponent struct {
	version  int
	opType   int
	contents []*bitComponent
	// for int literal
	value int
}

func (c bitComponent) sumVersions() int {
	sum := c.version

	for _, sub := range c.contents {
		sum += sub.sumVersions()
	}

	return sum
}

func (c bitComponent) evaluate() int {
	switch c.opType {
	case 0:
		sum := 0
		for _, sub := range c.contents {
			sum += sub.evaluate()
		}
		return sum

	case 1:
		prod := 1
		for _, sub := range c.contents {
			prod *= sub.evaluate()
		}
		return prod

	case 2:
		min := c.contents[0].evaluate()
		for i := 1; i < len(c.contents); i++ {
			v := c.contents[i].evaluate()
			if v < min {
				min = v
			}
		}
		return min

	case 3:
		max := c.contents[0].evaluate()
		for i := 1; i < len(c.contents); i++ {
			v := c.contents[i].evaluate()
			if v > max {
				max = v
			}
		}
		return max

	case 4:
		return c.value

	case 5:
		first := c.contents[0].evaluate()
		second := c.contents[1].evaluate()
		if first > second {
			return 1
		}
		return 0

	case 6:
		first := c.contents[0].evaluate()
		second := c.contents[1].evaluate()
		if first < second {
			return 1
		}
		return 0

	case 7:
		first := c.contents[0].evaluate()
		second := c.contents[1].evaluate()
		if first == second {
			return 1
		}
		return 0

	default:
		Fail("oops")
	}

	return -1
}

func (b bitComponent) print(level int) {
	fmt.Print(strings.Repeat("  ", level))
	switch b.opType {
	case 0:
		fmt.Println("+")
	case 1:
		fmt.Println("*")
	case 2:
		fmt.Println("min")
	case 3:
		fmt.Println("max")
	case 4:
		fmt.Println(b.value)
	case 5:
		fmt.Println(">")
	case 6:
		fmt.Println("<")
	case 7:
		fmt.Println("==")
	}

	for _, sub := range b.contents {
		sub.print(level + 1)
	}
}

var _ = Describe("16", func() {
	DescribeTable("example version sums", func(in string, sum int) {
		comp, _ := parseBitComponent(hexToBits(in))

		Expect(comp.sumVersions()).To(Equal(sum))
	},
		Entry("1", "8A004A801A8002F478", 16),
		Entry("2", "620080001611562C8802118E34", 12),
		Entry("3", "C0015000016115A2E0802F182340", 23),
		Entry("4", "A0016C880162017C3686B18A3D4780", 31),
	)

	It("does part A", func() {
		bs, err := ioutil.ReadFile("input16")
		Expect(err).NotTo(HaveOccurred())
		hexStr := strings.TrimSpace(string(bs))
		bitStr := hexToBits(hexStr)
		component, _ := parseBitComponent(bitStr)
		Expect(component.sumVersions()).To(Equal(917))
	})

	DescribeTable("example evaluations", func(in string, res int) {
		comp, _ := parseBitComponent(hexToBits(in))

		Expect(comp.evaluate()).To(Equal(res))
	},
		Entry("1", "C200B40A82", 3),
		Entry("2", "04005AC33890", 54),
		Entry("3", "880086C3E88112", 7),
		Entry("4", "CE00C43D881120", 9),
		Entry("5", "D8005AC2A8F0", 1),
		Entry("6", "F600BC2D8F", 0),
		Entry("7", "9C005AC2F8F0", 0),
		Entry("8", "9C0141080250320F1802104A08", 1),
	)

	It("does part B", func() {
		bs, err := ioutil.ReadFile("input16")
		Expect(err).NotTo(HaveOccurred())
		hexStr := strings.TrimSpace(string(bs))
		bitStr := hexToBits(hexStr)
		component, _ := parseBitComponent(bitStr)
		Expect(component.evaluate()).To(Equal(2536453523344))
	})
})

func hexToBits(hexStr string) string {
	var bitBuilder strings.Builder
	for _, r := range hexStr {
		n, err := strconv.ParseInt(string(r), 16, 8)
		Expect(err).NotTo(HaveOccurred())
		bits := fmt.Sprintf("%04b", n)
		bitBuilder.WriteString(bits)
	}

	return bitBuilder.String()
}

func parseBitComponent(s string) (*bitComponent, string) {
	compVersion := binToInt(s[0:3])
	compType := binToInt(s[3:6])

	if compType == 4 {
		rest := s[6:]
		cont := true
		intBits := ""
		for cont {
			if rest[0] == '0' {
				cont = false
			}
			intBits += rest[1:5]
			rest = rest[5:]
		}

		return &bitComponent{
			version: compVersion,
			opType:  4,
			value:   binToInt(intBits),
		}, rest
	}

	lenType := s[6:7]

	var subComponents []*bitComponent
	var restString string

	switch lenType {
	case "0":
		subPacketLen := binToInt(s[7:22])
		rest := s[22 : 22+subPacketLen]
		for len(rest) > 0 {
			subComponent, remaining := parseBitComponent(rest)
			rest = remaining
			subComponents = append(subComponents, subComponent)
		}
		restString = s[22+subPacketLen:]

	case "1":
		numSubPacketsStr := s[7:18]
		numSubPackets := binToInt(numSubPacketsStr)
		rest := s[18:]
		for i := 0; i < numSubPackets; i++ {
			subComponent, remaining := parseBitComponent(rest)
			subComponents = append(subComponents, subComponent)
			rest = remaining
		}
		restString = rest
	}

	return &bitComponent{
		version:  compVersion,
		opType:   compType,
		contents: subComponents,
	}, restString
}

func binToInt(b string) int {
	n, err := strconv.ParseInt(b, 2, 64)
	Expect(err).NotTo(HaveOccurred())

	return int(n)
}
