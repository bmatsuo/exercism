package etl

import "strings"

// Cardinality of the Scrabble alphabet.  If we are not restricted to latin
// characters this will be wrong.  But its only used to reduce allocations.
const alphabetSize = 26

func Transform(legacy map[int][]string) map[string]int {
	m := make(map[string]int, alphabetSize)
	for score := range legacy {
		for _, tile := range legacy[score] {
			tile = strings.ToLower(tile)
			m[tile] = score
		}
	}
	return m
}
