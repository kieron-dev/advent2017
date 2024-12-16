package twenty24

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type block struct {
	idx    int
	id     int
	empty  bool
	length int
	prev   *block
	next   *block
}

func newFile(idx, id, length int) *block {
	return &block{
		idx:    idx,
		id:     id,
		empty:  false,
		length: length,
	}
}

func newGap(idx, length int) *block {
	return &block{
		idx:    idx,
		empty:  true,
		length: length,
	}
}

func (b *block) insertNext(next *block) {
	oldNext := b.next
	b.next = next
	next.next = oldNext
	next.prev = b
	if oldNext != nil {
		oldNext.prev = next
	}
}

func (b *block) checksum(pos int) int {
	if b.empty {
		return 0
	}
	return b.id * b.length * (2*pos + b.length - 1) / 2
}

func (b *block) String() string {
	c := "."
	if !b.empty {
		c = strconv.Itoa(b.id)
	}
	return strings.Repeat(c+",", b.length)
}

func (b *block) Map(fn func(*block)) {
	it := b
	for it != nil {
		fn(it)
		it = it.next
	}
}

func (b *block) unlink() {
	oldLeft := b.prev
	oldRight := b.next
	if oldLeft != nil {
		oldLeft.next = oldRight
	}
	if oldRight != nil {
		oldRight.prev = oldLeft
	}
	b.prev = nil
	b.next = nil
	b.id = 0
}

func Test09A(t *testing.T) {
	f, err := os.Open("input09")
	assert.NoError(t, err)
	defer f.Close()

	bs, err := io.ReadAll(f)
	assert.NoError(t, err)

	// bs = []byte("12345")
	// bs = []byte("2333133121414131402")
	// bs = []byte("151010")

	var start *block
	var last *block
	for i, c := range bs {
		if c > '9' || c < '0' {
			continue
		}
		l := int(c - '0')
		var b *block
		if i%2 == 1 {
			b = newGap(i, l)
		} else {
			b = newFile(i, i/2, l)
		}
		if start == nil {
			start = b
		}
		if last != nil {
			last.insertNext(b)
		}
		last = b
	}
	end := last
	fmt.Printf("end = %+v\n", end)

	printMap := false
	if printMap {
		start.Map(func(b *block) {
			fmt.Printf("%s", b)
		})
		fmt.Println()
	}

	beforeBlocks := 0
	start.Map(func(b *block) {
		if !b.empty {
			beforeBlocks += b.length
		}
	})

	compact(t, start, end)

	afterBlocks := 0
	start.Map(func(b *block) {
		if !b.empty {
			afterBlocks += b.length
		}
	})

	assert.Equal(t, beforeBlocks, afterBlocks)

	if printMap {
		start.Map(func(b *block) {
			fmt.Printf("%s", b)
		})
		fmt.Println()
	}

	checkSum := 0
	i := 0
	start.Map(func(b *block) {
		checkSum += b.checksum(i)
		i += b.length
	})
	assert.Equal(t, 6370402949053, checkSum)
}

func Test09B(t *testing.T) {
	f, err := os.Open("input09")
	assert.NoError(t, err)
	defer f.Close()

	bs, err := io.ReadAll(f)
	assert.NoError(t, err)

	// bs = []byte("12345")
	// bs = []byte("2333133121414131402")
	// bs = []byte("151010")

	var start *block
	var last *block
	for i, c := range bs {
		if c > '9' || c < '0' {
			continue
		}
		l := int(c - '0')
		var b *block
		if i%2 == 1 {
			b = newGap(i, l)
		} else {
			b = newFile(i, i/2, l)
		}
		if start == nil {
			start = b
		}
		if last != nil {
			last.insertNext(b)
		}
		last = b
	}
	end := last
	fmt.Printf("end = %+v\n", end)

	printMap := false
	if printMap {
		start.Map(func(b *block) {
			fmt.Printf("%s", b)
		})
		fmt.Println()
	}

	beforeBlocks := 0
	start.Map(func(b *block) {
		if !b.empty {
			beforeBlocks += b.length
		}
	})

	compactB(t, start, end)

	afterBlocks := 0
	start.Map(func(b *block) {
		if !b.empty {
			afterBlocks += b.length
		}
	})

	assert.Equal(t, beforeBlocks, afterBlocks)

	if printMap {
		start.Map(func(b *block) {
			fmt.Printf("%s", b)
		})
		fmt.Println()
	}

	checkSum := 0
	i := 0
	start.Map(func(b *block) {
		checkSum += b.checksum(i)
		i += b.length
	})
	fmt.Printf("checkSum = %+v\n", checkSum)

	assert.Equal(t, 6398096697992, checkSum)
}

func compact(t *testing.T, start, end *block) {
	left := start
	right := end

	for left.idx < right.idx {
		for left != nil && !left.empty {
			left = left.next
		}
		for right != nil && right.empty {
			right = right.prev
		}
		if left == nil || right == nil {
			break
		}
		assert.True(t, left.empty)
		assert.False(t, right.empty)

		if left.length == 0 {
			nextLeft := left.next
			left.unlink()
			left = nextLeft
			continue
		}

		if right.length == 0 {
			nextRight := right.prev
			right.unlink()
			right = nextRight
			continue
		}

		if left.length <= right.length {
			left.empty = false
			left.id = right.id
			right.length -= left.length
			left = left.next
		} else {
			b := newGap(left.idx, left.length-right.length)
			left.insertNext(b)
			left.length = right.length
			left.id = right.id
			left.empty = false
			right.length = 0
		}

		if right.length == 0 {
			nextRight := right.prev
			right.unlink()
			right = nextRight
		}
	}
}

func compactB(t *testing.T, start, end *block) {
	for end.empty {
		end = end.prev
	}
	cur := end.id

	for cur >= 0 {
		left := start
		right := end

		for right != nil && (right.empty || right.id != cur) {
			right = right.prev
		}
		cur--
		for left != nil && (!left.empty || left.length < right.length) {
			left = left.next
		}
		if left == nil || right == nil || left.idx > right.idx {
			continue
		}
		assert.True(t, left.empty)
		assert.False(t, right.empty)

		// if left.length == 0 {
		// 	nextLeft := left.next
		// 	left.unlink()
		// 	left = nextLeft
		// 	continue
		// }

		// if right.length == 0 {
		// 	nextRight := right.prev
		// 	right.unlink()
		// 	right = nextRight
		// 	continue
		// }

		if left.length > right.length {
			b := newGap(left.idx, left.length-right.length)
			left.insertNext(b)
		}
		left.length = right.length
		left.id = right.id
		left.empty = false
		right.empty = true

	}
}
