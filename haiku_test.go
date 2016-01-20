package haiku

import (
	"bytes"
	"crypto/rand"
	"fmt"
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

func TestRnd(t *testing.T) {
	// Panic test
	randomSource = bytes.NewReader(nil)
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected the generator to panic with an invalid source")
			}
		}()
		fmt.Println(rnd(1))
	}()

	randomSource = rand.Reader
	// Check that two consecutive random numbers are not the same
	if r1, r2 := rnd(1000), rnd(1000); r1 == r2 {
		t.Errorf(
			"The number %d has been randomly generated two times in a row",
			r1,
		)
	}
}

func BenchmarkRnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rnd(1000)
	}
}
