package recursion

// Recursion function
func Recursion(a int, b int) int {
	if b == 0 {
		return 0
	}
	return Recursion(b, a%b)
}

// WithoutRecursion function
func WithoutRecursion(a int, b int) int {
	var val int
	val = a % b

	if val != 0 {
		val = a % val
	}

	return val
}
