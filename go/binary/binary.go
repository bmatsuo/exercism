package binary

// ToDecimal converts a string of '0'/'1' characters and returns its
// representation as an integer.  if an unexpected character is encountered, 0
// is retured.
func ToDecimal(bstr string) int {
	var x int
	for i := range bstr {
		switch bstr[i] {
		case '0':
			x <<= 1
		case '1':
			x = (x << 1) | 1
		default:
			return 0
		}
	}
	return x
}
