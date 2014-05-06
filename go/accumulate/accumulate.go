package accumulate

// Accumulate returns a slice containing the result of applying tr to each
// element of strs.
func Accumulate(strs []string, tr func(string) string) []string {
	if len(strs) == 0 {
		return nil
	}
	trs := make([]string, len(strs))
	for i := range strs {
		trs[i] = tr(strs[i])
	}
	return trs
}
