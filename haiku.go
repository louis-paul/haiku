// Package haiku offers a random, Heroku-like name generator. The generated
// names are of the form "lively-flower-9136".
package haiku

import (
	"fmt"
	"math/rand"
)

const (
	// minSuffix is the minimum number (comprised) at the end of the haiku
	minSuffix = 1000
	// maxSuffix is the maximum number (comprised) at the end of the haiku
	maxSuffix = 9999
)

var (
	// adjectives is the list of adjectives to use as the first part of the
	// haiku
	adjectives = []string{
		"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark",
		"summer", "icy", "delicate", "quiet", "white", "cool", "spring",
		"winter", "patient", "twilight", "dawn", "crimson", "wispy",
		"weathered", "blue", "billowing", "broken", "cold", "damp", "falling",
		"frosty", "green", "long", "late", "lingering", "bold", "little",
		"morning", "muddy", "old", "red", "rough", "still", "small",
		"sparkling", "throbbing", "shy", "wandering", "withered", "wild",
		"black", "young", "holy", "solitary", "fragrant", "aged", "snowy",
		"proud", "floral", "restless", "divine", "polished", "ancient",
		"purple", "lively", "nameless",
	}

	// nouns is the list of nouns to use in the second part of the haiku
	nouns = []string{
		"waterfall", "river", "breeze", "moon", "rain", "wind", "sea",
		"morning", "snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn",
		"glitter", "forest", "hill", "cloud", "meadow", "sun", "glade", "bird",
		"brook", "butterfly", "bush", "dew", "dust", "field", "fire", "flower",
		"firefly", "feather", "grass", "haze", "mountain", "night", "pond",
		"darkness", "snowflake", "silence", "sound", "sky", "shape", "surf",
		"thunder", "violet", "water", "wildflower", "wave", "water",
		"resonance", "sun", "wood", "dream", "cherry", "tree", "fog", "frost",
		"voice", "paper", "frog", "smoke", "star",
	}
)

// Haiku returns a name of the form "lively-flower-9136"
func Haiku() string {
	adjective := adjectives[rand.Int31n(int32(len(adjectives)-1))]
	noun := nouns[rand.Int31n(int32(len(nouns)-1))]
	number := minSuffix + rand.Int31n(maxSuffix-minSuffix)
	return fmt.Sprintf(
		"%s-%s-%d",
		adjective,
		noun,
		number,
	)
}
