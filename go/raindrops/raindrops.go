package raindrops

import (
	"strconv"
)

// Convert translates integers into randrop-speak.  The raindrop-speak
// specification can be found at
// http://exercism.io/exercises/go/raindrops/readme
func Convert(x int) string {
	var rain string
	for _, tr := range factorTranslations {
		if x%tr.prime == 0 {
			rain += tr.text
		}
	}
	if rain == "" {
		return strconv.Itoa(x)
	}
	return rain
}

// raindrop-speak prime factor translations
var factorTranslations = []struct {
	prime int
	text  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}
