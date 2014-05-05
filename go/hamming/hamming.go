package hamming

// Distance computes the Hamming distance between homologous DNA sequences h1
// and h2.  If h1 and h2 differ in length they will only be inspected upto the
// length of the shortest sequence.
func Distance(h1, h2 string) int {
	var dist int
	n, m := len(h1), len(h2)
	if m < n {
		n = m
	}
	for i := 0; i < n; i++ {
		if h1[i] != h2[i] {
			dist++
		}
	}
	return dist
}
