package accumulate

// Accumulate : Accumulates the given strings according to the function
func Accumulate(given []string, f func(string) string) []string {
	ans := make([]string, len(given))
	for i, str := range given {
		ans[i] = f(str)
	}
	return ans
}
