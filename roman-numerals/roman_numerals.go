package romannumerals

import (
	"errors"
)

// ToRomanNumeral : Converts to roman numeral
func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "0", errors.New("Invalid value for converting to Roman")
	}
	intToRomanMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 10, 9, 5, 4, 1}
	roman := ""
	for _, k := range keys {
		if num/k > 0 {
			count := num / k
			num = num % k
			for count > 0 {
				roman = roman + intToRomanMap[k]
				count--
			}
		}
		if num == 0 {
			break
		}
	}
	return roman, nil
}
