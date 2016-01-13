// Package seq implements functions for well-known sequences like Fibonacci.
package seq

// Fib returns nth (from 0th) Fibonaccci number.
func Fib(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}
