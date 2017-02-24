package fib

// Fib : generates a Fibonacci series and returns it in a slice
func Fib(n int) []int {
	var p []int
	first, second := 0, 1
	for i := 0; i < n; i++ {
		if i < 2 {
			p = append(p, i)
		} else {
			current := first + second
			first = second
			second = current
			p = append(p, current)
		}
	}
	return p
}
