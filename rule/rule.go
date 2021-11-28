package rule

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Rule ...
type Rule func() []rune

// String ...
func String(s string) Rule {
	return func() []rune {
		return []rune(s)
	}
}

// Mixin creates a rule that executes multiple rules.
func Mixin(rules ...Rule) Rule {
	return func() []rune {
		var out = make([]rune, 0)
		for _, rule := range rules {
			out = append(out, rule()...)
		}
		return out
	}
}
