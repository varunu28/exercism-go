package etl

import (
	"strings"
)

// Transform : Transforms to new scrabble score format
func Transform(inp map[int][]string) map[string]int {
	ans := make(map[string]int)
	for k, v := range inp {
		for _, letter := range v {
			ans[strings.ToLower(letter)] = k
		}
	}
	return ans
}
