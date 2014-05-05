package grains

import "math"

// Square returns the number of grains of wheat on square i of the chessboard.
func Square(i int) uint64 {
	if i < 1 {
		panic("invalid square")
	}
	return 1 << uint(i-1)
}

// TotalGrains is the number of grains on the chessboard.  The number of grains
// on each square is a power of two (binary representation: 00...00100...00).
// The sum of these 64 unique values is a 64 bit unsigned integer with all bits
// set to one.
const TotalGrains = math.MaxUint64

// Total returns the number of grains on the chessboard. See TotalGrains.
func Total() uint64 {
	return TotalGrains
}
