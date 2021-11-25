package rule

import (
	"math/rand"
	"strconv"
)

// RandomAlphabet ...
func RandomAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			switch rand.Intn(2) {
			case 0:
				out[i] = 'a' + rune(rand.Intn(26))
			case 1:
				out[i] = 'A' + rune(rand.Intn(26))
			}
		}
		return out
	}
}

// RandomUppercaseAlphabet ...
func RandomUppercaseAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = 'A' + rune(rand.Intn(26))
		}
		return out
	}
}

// RandomLowercaseAlphabet ...
func RandomLowercaseAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = 'a' + rune(rand.Intn(26))
		}
		return out
	}
}

// RandomNumeric ...
func RandomNumeric(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = []rune(strconv.Itoa(rand.Intn(10)))[0]
		}
		return out
	}
}

// RandomAlphabetAndNumeric ...
func RandomAlphabetAndNumeric(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			switch rand.Intn(3) {
			case 0:
				out[i] = 'a' + rune(rand.Intn(26))
			case 1:
				out[i] = 'A' + rune(rand.Intn(26))
			case 2:
				out[i] = []rune(strconv.Itoa(rand.Intn(10)))[0]
			}
		}
		return out
	}
}
