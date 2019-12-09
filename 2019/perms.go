package advent2019

type Perms struct{}

func (p Perms) All(elements []int64) [][]int64 {
	var res [][]int64
	generate(len(elements), elements, &res)
	return res
}

func generate(k int, elements []int64, out *[][]int64) {
	if k == 1 {
		p := make([]int64, len(elements))
		copy(p, elements)
		*out = append(*out, p)
		return
	}

	generate(k-1, elements, out)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			elements[i], elements[k-1] = elements[k-1], elements[i]
		} else {
			elements[0], elements[k-1] = elements[k-1], elements[0]
		}
		generate(k-1, elements, out)
	}
}
