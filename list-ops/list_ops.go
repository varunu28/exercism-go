package listops

// IntList type definition of int slice
type IntList []int

// predFunc  function that returns a boolean value given an int value
type predFunc func(n int) bool

// unaryFunc function that returns an int value given an int value
type unaryFunc func(n int) int

// binFunc function that returns an int value given pair of int values
type binFunc func(a, b int) int

// Append appends all items in the second list to the end of the first list
func (l IntList) Append(inp IntList) IntList {
	ans := make(IntList, l.Length()+inp.Length())
	count := l.Length()
	for i, item := range l {
		ans[i] = item
	}
	for i, item := range inp {
		ans[count+i] = item
	}
	return ans
}

// Concat given a series of lists, combine all items in all lists into one flattened list
func (l IntList) Concat(inps []IntList) IntList {
	for _, inp := range inps {
		l = l.Append(inp)
	}
	return l
}

// Filter given a predicate and a list, return the list of all items for which predicate(item) is True
func (l IntList) Filter(f predFunc) IntList {
	ans := make(IntList, 0, l.Length())
	idx := 0
	for _, item := range l {
		if f(item) {
			ans = ans[:idx+1]
			ans[idx] = item
			idx++
		}
	}
	return ans
}

// Length given a list, return the total number of items within it
func (l IntList) Length() int {
	count := 0
	for range l {
		count++
	}
	return count
}

// Map given a function and a list, return the list of the results of applying function(item) on all items
func (l IntList) Map(f unaryFunc) IntList {
	for i, item := range l {
		l[i] = f(item)
	}
	return l
}

// Foldr given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the right using function(item, accumulator)
func (l IntList) Foldr(f binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		initial = f(l[i], initial)
	}
	return initial
}

// Foldl given a function, a list, and initial accumulator, fold (reduce) each item into the accumulator from the left using function(accumulator, item)
func (l IntList) Foldl(f binFunc, initial int) int {
	for _, item := range l {
		initial = f(initial, item)
	}
	return initial
}

// Reverse given a list, return a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < j; i++ {
		l[i], l[j] = l[j], l[i]
		j--
	}
	return l
}
