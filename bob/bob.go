// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var isQuestion = false
	var isYelledCount = 0
	var charCount = 0
	var nonSpaceCount = 0
	var questionIdx = 0
	for i, c := range remark {
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			continue
		}
		nonSpaceCount++
		if unicode.IsLetter(c) {
			charCount++
			if unicode.IsUpper(c) {
				isYelledCount++
			}
		}
		if c == '?' {
			isQuestion = true
			questionIdx = i + 1
			break
		}
	}
	for questionIdx < len(remark) {
		if remark[questionIdx] != ' ' {
			isQuestion = false
			break
		}
		questionIdx++
	}
	var isYelled = isYelledCount == charCount && isYelledCount > 0
	if nonSpaceCount == 0 {
		return "Fine. Be that way!"
	} else if isQuestion && isYelled {
		return "Calm down, I know what I'm doing!"
	} else if isYelled {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
