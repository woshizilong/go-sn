package rule

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	got := String("test")()
	want := []rune("test")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestMixin(t *testing.T) {
	got := Mixin(
		String("hello"),
		String("world"),
	)()
	want := []rune("helloworld")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
