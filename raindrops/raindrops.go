package raindrops

import "strconv"

// Convert : Converts to raindrop sound
func Convert(num int) string {
	res := ""
	if num%3 == 0 {
		res += "Pling"
	}
	if num%5 == 0 {
		res += "Plang"
	}
	if num%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return strconv.Itoa(num)
	}
	return res
}
