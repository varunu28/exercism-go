package grains

import (
	"errors"
)

// Square returns number of grains on the square
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid value for square")
	}
	return uint64(1 << uint(n-1)), nil
}

// Total returns total number of grains on board
func Total() uint64 {
	return uint64((1 << 64) - 1)
}
