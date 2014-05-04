package raindrops

import (
	"bytes"
	"strconv"
)

// Convert translates integers into randrop-speak.  The randrop-speak
// specification can be found at
// http://exercism.io/exercises/go/raindrops/readme
func Convert(x int) string {
	var buf bytes.Buffer
	var primeFound bool
	for _, tr := range factorTranslations {
		if x%tr.prime == 0 {
			primeFound = true
			buf.Write(tr.text)
		}
	}
	if !primeFound {
		return strconv.Itoa(x)
	}
	return buf.String()
}

// raindrop-speak prime factor translations
var factorTranslations = []struct {
	prime int
	text  []byte
}{
	{3, []byte("Pling")},
	{5, []byte("Plang")},
	{7, []byte("Plong")},
}
