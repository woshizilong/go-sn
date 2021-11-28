package rule

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestRandomAlphabet(t *testing.T) {
	rand.Seed(0)
	rule := RandomAlphabet(4)
	table := [][]rune{
		[]rune("uYIZ"),
		[]rune("abEa"),
		[]rune("ScMg"),
		[]rune("CyWE"),
		[]rune("WhiG"),
		[]rune("MaHu"),
		[]rune("hceE"),
		[]rune("CepM"),
		[]rune("pOTQ"),
		[]rune("plYJ"),
		[]rune("oxfk"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestRandomUppercaseAlphabet(t *testing.T) {
	rand.Seed(0)
	rule := RandomUppercaseAlphabet(4)
	table := [][]rune{
		[]rune("CUBY"),
		[]rune("HIZZ"),
		[]rune("KAKB"),
		[]rune("LEEA"),
		[]rune("NSOC"),
		[]rune("PMIG"),
		[]rune("NCKY"),
		[]rune("RWXE"),
		[]rune("VWSH"),
		[]rune("OIPG"),
		[]rune("ZMGA"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}

func TestRandomLowercaseAlphabet(t *testing.T) {
	rand.Seed(0)
	rule := RandomLowercaseAlphabet(4)
	table := [][]rune{
		[]rune("cuby"),
		[]rune("hizz"),
		[]rune("kakb"),
		[]rune("leea"),
		[]rune("nsoc"),
		[]rune("pmig"),
		[]rune("ncky"),
		[]rune("rwxe"),
		[]rune("vwsh"),
		[]rune("oipg"),
		[]rune("zmga"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}

func TestRandomNumeric(t *testing.T) {
	rand.Seed(0)
	rule := RandomNumeric(4)
	table := [][]rune{
		[]rune("4436"),
		[]rune("5677"),
		[]rune("8887"),
		[]rune("9826"),
		[]rune("1004"),
		[]rune("5860"),
		[]rune("1620"),
		[]rune("1278"),
		[]rune("5483"),
		[]rune("0094"),
		[]rune("7468"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}
func TestRandomAlphabetAndNumeric(t *testing.T) {
	rand.Seed(0)
	rule := RandomAlphabetAndNumeric(4)
	table := [][]rune{
		[]rune("uY6Z"),
		[]rune("ab8a"),
		[]rune("0C80"),
		[]rune("6YWe"),
		[]rune("4304"),
		[]rune("MaHu"),
		[]rune("h8eE"),
		[]rune("c4pm"),
		[]rune("PO1q"),
		[]rune("7LYj"),
		[]rune("o1F6"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}

func TestRandomList(t *testing.T) {
	rand.Seed(0)
	rule := RandomList(
		[]string{"Dog", "Cat", "Fox", "Rat"},
	)
	table := [][]rune{
		[]rune("Fox"),
		[]rune("Fox"),
		[]rune("Cat"),
		[]rune("Fox"),
		[]rune("Rat"),
		[]rune("Dog"),
		[]rune("Rat"),
		[]rune("Cat"),
		[]rune("Dog"),
		[]rune("Dog"),
		[]rune("Dog"),
	}
	for i := 0; i < 11; i++ {
		if got, want := rule(), table[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}
