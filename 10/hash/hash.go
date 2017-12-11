package hash

import "fmt"

func Get(list []int, pos int) int {
	l := len(list)
	return list[pos%l]
}

func Set(list []int, pos, val int) {
	list[pos%len(list)] = val
}

func Reverse(list []int, start, end int) {
	for end > start {
		sVal := Get(list, start)
		eVal := Get(list, end)
		Set(list, end, sVal)
		Set(list, start, eVal)
		start++
		end--
	}
}

func Compute(instructions []int, listLen int) int {
	list := []int{}
	for i := 0; i < listLen; i++ {
		list = append(list, i)
	}

	pos := 0
	skip := 0
	for _, n := range instructions {
		Reverse(list, pos, pos+n-1)
		pos += n + skip
		skip++
	}
	return list[0] * list[1]
}

func Compute2(instructions []byte, listLen int) string {
	list := []int{}
	for i := 0; i < listLen; i++ {
		list = append(list, i)
	}

	pos := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, b := range instructions {
			n := int(b)
			Reverse(list, pos, pos+n-1)
			pos += n + skip
			skip++
		}
	}
	dense := xorIt(list)
	return toHex(dense)
}

func xorIt(list []int) []int {
	sections := len(list) / 16
	ret := []int{}

	for i := 0; i < sections; i++ {
		part := 0
		for j := 0; j < 16; j++ {
			part ^= list[j+i*16]
		}
		ret = append(ret, part)
	}
	return ret
}

func toHex(list []int) string {
	ret := ""
	for _, i := range list {
		ret += fmt.Sprintf("%0x", i)
	}
	return ret
}
