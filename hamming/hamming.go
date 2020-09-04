package hamming

import (
	"fmt"
)

// Distance Returns the hamming distance between 2 DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Mismatch in length of 2 strands")
	}
	mismatchCount := 0
	for i := range a {
		if a[i] != b[i] {
			mismatchCount = mismatchCount + 1
		}
	}
	return mismatchCount, nil
}
