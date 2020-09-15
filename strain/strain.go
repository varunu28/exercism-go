package strain

// Ints : Slice of integers
type Ints []int

// Lists : List of slices
type Lists [][]int

// Strings : Slice of strings
type Strings []string

// Keep : Filters only values that evaluate to true for Ints
func (i Ints) Keep(f func(int) bool) (o Ints) {
	for _, num := range i {
		if f(num) {
			o = append(o, num)
		}
	}
	return
}

// Discard : Filters only values that evaluate to false for Ints
func (i Ints) Discard(f func(int) bool) (o Ints) {
	for _, num := range i {
		if !f(num) {
			o = append(o, num)
		}
	}
	return
}

// Keep : Filters only values that evaluate to true for Lists
func (l Lists) Keep(f func([]int) bool) (o Lists) {
	for _, lst := range l {
		if f(lst) {
			o = append(o, lst)
		}
	}
	return
}

// Keep : Filters only values that evaluate to true for Strings
func (s Strings) Keep(f func(string) bool) (o Strings) {
	for _, str := range s {
		if f(str) {
			o = append(o, str)
		}
	}
	return
}
