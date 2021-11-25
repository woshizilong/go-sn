package rule

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// Numbering ...
func Numbering() Rule {
	var n int64 = 0
	return func() []rune {
		out := []rune(strconv.FormatInt(n, 10))
		atomic.AddInt64(&n, 1)
		return out
	}
}

// NumberingWithPadding ...
func NumberingWithPadding(pad int) Rule {
	var n int64 = 0
	var format = "%0" + strconv.Itoa(pad) + "d"
	return func() []rune {
		out := []rune(fmt.Sprintf(format, n))
		atomic.AddInt64(&n, 1)
		return out
	}
}
