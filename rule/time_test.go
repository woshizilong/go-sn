package rule

import (
	"reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeNow = func() time.Time {
		return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	}
	got := Time(time.RFC3339Nano)()
	want := []rune("2021-01-01T01:01:01.000000001Z")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestUnixXXXX(t *testing.T) {
	timeNow = func() time.Time {
		return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	}

	t.Run("case=UnixNano", func(t *testing.T) {
		got := UnixNano()()
		want := []rune("1609462861000000001")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=UnixMicro", func(t *testing.T) {
		got := UnixMicro()()
		want := []rune("1609462861000000")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=UnixMilli", func(t *testing.T) {
		got := UnixMilli()()
		want := []rune("1609462861000")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
