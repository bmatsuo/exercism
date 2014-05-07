package strand

import "strings"

// ToRNA returns the RNA sequence transcribed from dna.  Runes not representing
// a nucleotide are left unchanged.
func ToRna(dna string) string {
	return strings.Map(trRNA, dna)
}

func trRNA(c rune) rune {
	switch c {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return c
	}
}
