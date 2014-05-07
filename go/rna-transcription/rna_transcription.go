package strand

import "strings"

// ToRNA returns the RNA sequence transcribed from dna.  Runes not representing
// a nucleotide are left unchanged.
func ToRna(input string) string {
	return trRNA.Replace(input)
}

var trRNA = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)
