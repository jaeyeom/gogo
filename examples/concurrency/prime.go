package concurrency

import (
	"fmt"

	"golang.org/x/net/context"
)

// Range returns a channel and sends ints
// (start, start+step, start+2*step, ...).
func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// FilterMultiple returns a IntPipe that filters multiple of n.
func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Print(prime, " ")
	}
	fmt.Println()
}
