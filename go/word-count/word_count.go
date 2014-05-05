package wc

import (
	"strings"
	"unicode"
)

// WordCount returns a historgram containing the number of occurrances of each
// word in text.  A word is a contiguous sequence of letters or numbers.
func WordCount(text string) Histogram {
	h := make(Histogram)
	for _, word := range strings.FieldsFunc(text, delim) {
		word = strings.ToLower(word)
		h[word]++
	}
	return h
}

func delim(c rune) bool {
	return !(unicode.IsLetter(c) || unicode.IsNumber(c))
}
