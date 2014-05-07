package atbash

import "strings"

// Atbash encodes ASCII text.  Non-alphanumeric characters are removed from
// cipher text.  Long strings have spaces inserted after every 5th character
// for readability.
func Atbash(text string) string {
	if text == "" {
		return ""
	}

	cs := make([]rune, 0, len(text))
	n := 0
	put := func(c rune) {
		if n%5 == 0 && n > 0 {
			cs = append(cs, ' ')
		}
		n++
		cs = append(cs, c)
	}

	text = strings.ToLower(text)
	for _, c := range text {
		switch {
		case c >= '0' && c <= '9':
			put(c)
		case c >= 'a' && c <= 'z':
			put(('z' - c) + 'a')
		}
	}

	return string(cs)
}
