package anagram

import (
	"sort"
	"strings"
)

// Detect returns elements of candidates which are anagrams of target.
// Anagrams are case insensitive and all letters in returned strings will be in
// lower case. A word is not an anagram of itself.
//
// BUG no text normalization or multi-rune symbol handling (only works with ascii).
func Detect(target string, candidates []string) []string {
	var a []string
	tnorm := strings.ToLower(target)
	tsorted := runeSorted(tnorm)
	n := len(target)
	for _, s := range candidates {
		if len(s) != n {
			continue
		}
		snorm := strings.ToLower(s)
		if snorm == tnorm {
			continue
		}
		if runeSorted(snorm) == tsorted {
			a = append(a, snorm)
		}
	}
	return a
}

func runeSorted(str string) string {
	rs := runes(str)
	sort.Sort(rs)
	return string(rs)
}

type runes []rune

func (rs runes) Len() int           { return len(rs) }
func (rs runes) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs runes) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
