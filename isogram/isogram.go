package isogram

import (
	"unicode"
)

// IsIsogram checks if the given string is a Isogram
func IsIsogram(s string) bool {
	seen := [26]bool{}
	for _, c := range s {
		if unicode.IsLetter(c) {
			key := int(unicode.ToLower(c)) - 97
			if seen[key] {
				return false
			}
			seen[key] = true
		}
	}
	return true
}
