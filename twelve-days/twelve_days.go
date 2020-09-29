package twelve

import (
	"fmt"
	"strings"
)

// DayToMessageMap a map from day to message item string
var DayToMessageMap = map[string]string{
	"first":    "a Partridge in a Pear Tree",
	"second":   "two Turtle Doves",
	"third":    "three French Hens",
	"fourth":   "four Calling Birds",
	"fifth":    "five Gold Rings",
	"sixth":    "six Geese-a-Laying",
	"seventh":  "seven Swans-a-Swimming",
	"eighth":   "eight Maids-a-Milking",
	"ninth":    "nine Ladies Dancing",
	"tenth":    "ten Lords-a-Leaping",
	"eleventh": "eleven Pipers Piping",
	"twelfth":  "twelve Drummers Drumming",
}

// DayToDayStr mapping between int day to string day
var DayToDayStr = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func getFormattedGiftString(day int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", DayToDayStr[day])
}

// Song returns the complete song
func Song() string {
	ans := make([]string, 12)
	for i := 1; i <= 12; i++ {
		ans[i-1] = Verse(i)
	}
	return strings.Join(ans, "\n")
}

// Verse returns verse on line number
func Verse(line int) string {
	ans := getFormattedGiftString(line)
	for i := line; i > 0; i-- {
		if i == 1 {
			if line == 1 {
				ans = ans + DayToMessageMap[DayToDayStr[i]] + "."
			} else {
				ans = ans + "and " + DayToMessageMap[DayToDayStr[i]] + "."
			}
		} else {
			ans = ans + DayToMessageMap[DayToDayStr[i]] + ", "
		}
	}
	return ans
}
