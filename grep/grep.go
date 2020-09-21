package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// FlagMatch struct definition for all flag values
type FlagMatch struct {
	printNum        bool
	printOnlyFile   bool
	caseInsensitive bool
	invertMatch     bool
	entireLineMatch bool
}

// Search performs grep operation
func Search(pattern string, flags []string, files []string) []string {
	ans := []string{}
	seen := map[string]bool{}
	multipleFiles := len(files) > 1
	for _, filename := range files {
		file, _ := os.Open(filename)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			ismatch, match := isMatch(pattern, flags, filename, scanner.Text(), lineNum, multipleFiles)
			if ismatch {
				if _, ok := seen[match]; !ok {
					ans = append(ans, match)
					seen[match] = true
				}
			}
			lineNum++
		}
	}
	return ans
}

func getFlagMatch(flags []string) FlagMatch {
	flagMatch := FlagMatch{
		printNum:        false,
		printOnlyFile:   false,
		caseInsensitive: false,
		invertMatch:     false,
		entireLineMatch: false,
	}
	for _, flag := range flags {
		if flag == "-n" {
			flagMatch.printNum = true
		} else if flag == "-l" {
			flagMatch.printOnlyFile = true
		} else if flag == "-i" {
			flagMatch.caseInsensitive = true
		} else if flag == "-v" {
			flagMatch.invertMatch = true
		} else if flag == "-x" {
			flagMatch.entireLineMatch = true
		}
	}
	return flagMatch
}

func isMatch(pattern string, flags []string, filename string, text string, lineNum int, multipleFiles bool) (bool, string) {
	flagMatch := getFlagMatch(flags)
	match := false
	matchStr := ""
	if flagMatch.entireLineMatch {
		if flagMatch.caseInsensitive {
			if strings.ToLower(pattern) == strings.ToLower(text) {
				match = true
			}
		} else {
			match = pattern == text
		}
	} else {
		if flagMatch.caseInsensitive {
			if strings.Contains(strings.ToLower(text), strings.ToLower(pattern)) {
				match = true
			}
		} else {
			if strings.Contains(text, pattern) {
				match = true
			}
		}
	}
	if flagMatch.invertMatch {
		match = !match
	}
	if match {
		if flagMatch.printOnlyFile {
			matchStr = filename
		} else if flagMatch.printNum {
			if multipleFiles {
				matchStr = filename + ":" + strconv.Itoa(lineNum) + ":" + text
			} else {
				matchStr = strconv.Itoa(lineNum) + ":" + text
			}
		} else {
			if multipleFiles {
				matchStr = filename + ":" + text
			} else {
				matchStr = text
			}
		}
	}
	return match, matchStr
}
