package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks if given string is a valid ISBN
func IsValidISBN(s string) bool {
	s = strings.ReplaceAll(s, "-", "")
	r := []rune(s)
	if len(r) != 10 {
		return false
	}
	sum, mul, checkDigit := 0, 10, 0
	if unicode.IsDigit(r[len(r)-1]) {
		checkDigit = int(r[len(r)-1] - '0')
	} else if r[len(r)-1] == 'X' {
		checkDigit = 10
	} else {
		return false
	}
	for i := 0; i < len(r)-1; i++ {
		if !unicode.IsDigit(r[i]) {
			return false
		}
		sum += mul * int(r[i]-'0')
		mul--
	}
	sum += mul * checkDigit
	return sum%11 == 0
}
