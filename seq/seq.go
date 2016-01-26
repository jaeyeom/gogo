// Package seq implements functions for well-known sequences like Fibonacci.
package seq

// FibNth returns nth (from 0th) Fibonaccci sequence number.
func FibNth(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}

// FibChan returns a channel that emits the Fibonacci sequence up to max.
func FibChan(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

// FibGen returns a func that returns Fibonacci sequence numbers for
// each call up to max. If the number is bigger than max, it returns
// -1.
func FibGen(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}
