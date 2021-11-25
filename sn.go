package sn

import (
	"strings"

	"github.com/kenkyu392/go-sn/rule"
)

const (
	// DefaultDelimiter is the delimiter character
	// used in the Generator created by NewGenerator.
	DefaultDelimiter = "-"
)

// Generator ...
type Generator struct {
	Delimiter string
	Rules     []rule.Rule
}

// NewGenerator ...
func NewGenerator(rules ...rule.Rule) *Generator {
	return &Generator{
		Delimiter: DefaultDelimiter,
		Rules:     rules,
	}
}

// String implements fmt.Stringer interface.
func (g *Generator) String() string {
	var elems = make([]string, len(g.Rules))
	for i, rule := range g.Rules {
		elems[i] = string(rule())
	}
	return strings.Join(elems, g.Delimiter)
}
