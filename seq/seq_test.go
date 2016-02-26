package seq

import "fmt"

func ExampleFibNth() {
	for i := 0; i < 8; i++ {
		fmt.Print(FibNth(i), ",")
	}
	// Output: 0,1,1,2,3,5,8,13,
}

func ExampleFibChan() {
	for fib := range FibChan(15) {
		fmt.Print(fib, ",")
	}
	// Output: 0,1,1,2,3,5,8,13,
}

func ExampleFibGen() {
	fib := FibGen(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
	// Output: 0,1,1,2,3,5,8,13,
}
