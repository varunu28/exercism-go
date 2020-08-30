// Package acronym contains the function Abbreviate to find abbreciation of a string
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate finds abbreciation of a string
func Abbreviate(s string) string {
	const InitialVal rune = '-'
	var ans string = ""
	var prefix = InitialVal
	for i, c := range s {
		_ = i
		if c == ' ' || c == '-' {
			if prefix != InitialVal {
				ans = ans + strings.ToUpper(string(prefix))
				prefix = InitialVal
			}
		} else if unicode.IsLetter(c) && prefix == InitialVal {
			prefix = c
		}
	}
	if prefix != InitialVal {
		ans = ans + strings.ToUpper(string(prefix))
	}
	return ans
}
