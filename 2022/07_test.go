package two022_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type file struct {
	isDir    bool
	name     string
	contents map[string]*file
	parent   *file
	size     int
}

func newDir(name string, parent *file) *file {
	d := &file{
		name:     name,
		isDir:    true,
		contents: map[string]*file{},
		parent:   parent,
	}

	if parent == nil {
		d.parent = d
	}

	return d
}

func (f *file) cd(d string) *file {
	if d == ".." {
		return f.parent
	}

	return f.contents[d]
}

func (f *file) addFile(name string, size int) {
	f.contents[name] = &file{
		name:   name,
		parent: f,
		size:   size,
	}
}

func (f *file) addDir(name string) {
	f.contents[name] = newDir(name, f)
}

func (f *file) Size(cb func(int)) int {
	size := 0
	for _, fd := range f.contents {
		if fd.isDir {
			size += fd.Size(cb)
		} else {
			size += fd.size
		}
	}

	cb(size)
	return size
}

var _ = Describe("2022/07", func() {
	It("does part A & B", func() {
		f, err := os.Open("input07")
		Expect(err).NotTo(HaveOccurred())

		curDir := newDir("/", nil)
		rootDir := curDir

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()

			if strings.HasPrefix(line, "$ cd") {
				var d string
				_, err := fmt.Sscanf(line, "$ cd %s", &d)
				Expect(err).NotTo(HaveOccurred())

				if d == "/" {
					curDir = rootDir
				} else {
					curDir = curDir.cd(d)
				}

				continue
			}

			if strings.HasPrefix(line, "$ ls") {
				continue
			}

			if strings.HasPrefix(line, "dir") {
				var name string
				_, err := fmt.Sscanf(line, "dir %s", &name)
				Expect(err).NotTo(HaveOccurred())
				curDir.addDir(name)

				continue
			}

			var name string
			var size int
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			Expect(err).NotTo(HaveOccurred())
			curDir.addFile(name, size)
		}

		t := 0
		used := rootDir.Size(func(s int) {
			if s <= 100000 {
				t += s
			}
		})

		Expect(t).To(Equal(1086293))

		toFree := 30000000 - (70000000 - used)
		smallest := 70000000
		rootDir.Size(func(s int) {
			if s >= toFree && s < smallest {
				smallest = s
			}
		})

		Expect(smallest).To(Equal(366028))
	})
})
