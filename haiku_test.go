package haiku

import (
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	haikuRegex := regexp.MustCompile(`^\w+-\w+-\d{4}$`)
	if h := Haiku(); !haikuRegex.MatchString(h) {
		t.Errorf("The haiku %s does not match the testing regex", h)
	}
}

func TestRandomness(t *testing.T) {
	// Generate two haikus sequentially to check that they are not the same
	if h1, h2 := Haiku(), Haiku(); h1 == h2 {
		t.Errorf("The haiku %s has been generated two times in a row", h1)
	}
}

func BenchmarkHaiku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Haiku()
	}
}
