package brackets

func process(str string) (bracketCount int, garbageChars int) {
	c := 0
	g := 0
	level := 0
	garbage := false
	ignoreNext := false

	for _, r := range str {
		if garbage {
			if ignoreNext {
				ignoreNext = false
				continue
			}
			if r == '!' {
				ignoreNext = true
				continue
			} else if r == '>' {
				garbage = false
				continue
			}
			g++
			continue
		}
		if r == '<' {
			garbage = true
			continue
		}
		if r == '{' {
			level++
			c += level
		} else if r == '}' {
			level--
		}
	}
	return c, g
}

func Count(str string) int {
	val, _ := process(str)
	return val
}

func Garbage(str string) int {
	_, val := process(str)
	return val
}
