package rule

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestNumbering(t *testing.T) {
	rule := Numbering()
	for i := 0; i < 11; i++ {
		got := rule()
		want := []rune(strconv.Itoa(i))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestNumberingWithPadding(t *testing.T) {
	rule := NumberingWithPadding(6)
	for i := 0; i < 11; i++ {
		got := rule()
		want := []rune(fmt.Sprintf("%06d", i))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}
