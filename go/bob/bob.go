package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to statement.
func Hey(statement string) string {
	if isSilence(statement) {
		return "Fine. Be that way!"
	}
	if isShouting(statement) {
		return "Woah, chill out!"
	}
	if isQuestion(statement) {
		return "Sure."
	}
	return "Whatever."
}

func isSilence(statement string) bool {
	return allString(statement, unicode.IsSpace)
}

func isQuestion(statement string) bool {
	return strings.HasSuffix(statement, "?")
}

func isShouting(statement string) bool {
	return containsFunc(statement, unicode.IsLetter) && allString(statement, isShoutingRune)
}

func isShoutingRune(c rune) bool {
	return unicode.IsUpper(c) || !unicode.IsLetter(c)
}

func containsFunc(s string, fn func(c rune) bool) bool {
	return strings.IndexFunc(s, fn) > -1
}

func allString(s string, fn func(c rune) bool) bool {
	return !containsFunc(s, func(c rune) bool { return !fn(c) })
}
