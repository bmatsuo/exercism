package romannumerals

const (
	M = 1000
	D = 500
	C = 100
	L = 50
	X = 10
	V = 5
	I = 1
)

// ToRomanNumeral returns n's representation as a roman numeral.  An empty
// string is returned if n is 0. Behavior for numbers larger than 3999
// (MMMCMXCIX) is not defined.
func ToRomanNumeral(n int) string {
	if n == 0 {
		return ""
	}
	if n >= 4*M {
		panic("too large; roman numerals are terrible")
	}
	if n >= M {
		return "M" + ToRomanNumeral(n-M)
	}
	if n >= M-C {
		return "CM" + ToRomanNumeral(n-M+C)
	}
	if n >= D {
		return "D" + ToRomanNumeral(n-D)
	}
	if n >= D-C {
		return "CD" + ToRomanNumeral(n-D+C)
	}
	if n >= C {
		return "C" + ToRomanNumeral(n-C)
	}
	if n >= C-X {
		return "XC" + ToRomanNumeral(n-C+X)
	}
	if n >= L {
		return "L" + ToRomanNumeral(n-L)
	}
	if n >= L-X {
		return "XL" + ToRomanNumeral(n-L+X)
	}
	if n >= X {
		return "X" + ToRomanNumeral(n-X)
	}
	if n >= X-I {
		return "IX" + ToRomanNumeral(n-X)
	}
	if n >= V {
		return "V" + ToRomanNumeral(n-V)
	}
	if n >= V-I {
		return "IV" + ToRomanNumeral(n-V+I)
	}
	if n >= I {
		return "I" + ToRomanNumeral(n-I)
	}
	return ""
}
