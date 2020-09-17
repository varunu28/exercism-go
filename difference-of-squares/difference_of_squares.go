package diffsquares

// SquareOfSum returns square of sum of n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns sum of squares of n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns difference between SquareOfSum and SumOfSquares of n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
