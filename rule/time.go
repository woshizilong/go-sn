package rule

import (
	"strconv"
	"time"
)

var timeNow = time.Now

// Time ...
func Time(layout string) Rule {
	return func() []rune {
		return []rune(timeNow().Format(layout))
	}
}

// UnixNano ...
func UnixNano() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixNano())))
	}
}

// UnixMicro ...
func UnixMicro() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMicro())))
	}
}

// UnixMilli ...
func UnixMilli() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMilli())))
	}
}
