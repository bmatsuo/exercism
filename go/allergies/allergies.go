package allergies

var scores = map[string]int{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns the set of known allergens contributing to a total score.
func Allergies(score int) []string {
	var allergens []string
	for allergen, ascore := range scores {
		if score&ascore == ascore {
			allergens = append(allergens, allergen)
		}
	}
	return allergens
}

// AllergicTo returns true if the target allergen is a component of the total
// score.
func AllergicTo(score int, target string) bool {
	tscore, ok := scores[target]
	return ok && score&tscore == tscore
}
