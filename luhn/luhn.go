package luhn

import (
	"unicode"
)

// Valid checks if a string is valid per the Luhn formula
func Valid(s string) bool {
	r := []rune(s)
	sum := 0
	counter := 0
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsDigit(r[i]) {
			num := int(r[i] - '0')
			if counter%2 == 1 {
				num <<= 1
				if num > 9 {
					num -= 9
				}
			}
			sum += num
			counter++
		} else if !unicode.IsSpace(r[i]) {
			return false
		}
	}
	return counter > 1 && sum%10 == 0
}
