package advent2019

import (
	"bufio"
	"io"
)

type FileReader struct{}

func (f FileReader) Each(file io.Reader, fn func(line string)) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}
