package proverb

import "fmt"

// Proverb generates proverb from given list of words
func Proverb(rhyme []string) []string {
	if rhyme == nil || len(rhyme) == 0 {
		return []string{}
	}
	n := len(rhyme)
	proverb := make([]string, n)
	for i, j := 0, 0; i < len(rhyme)-1; i++ {
		proverb[j] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		j++
	}
	proverb[n-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return proverb
}
