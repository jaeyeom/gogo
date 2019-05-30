package calc

import "fmt"

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})
	exs := []string{
		"5",
		"1 + 2",
		"1 - 2 - 4",
		"( 3 - 2 ** 3 ) * ( -2 )",
		"3 * ( 3 + 1 * 3 ) / ( -2 )",
		"3 * ( ( 3 + 1 ) * 3 ) / 2",
		"1 + 2 ** 10 * 2",
		"2 ** 3 mod 3",
		"2 ** 2 ** 3",
	}
	for _, ex := range exs {
		fmt.Printf("%s = %d\n", ex, eval(ex))
	}
	// Output:
	// 5 = 5
	// 1 + 2 = 3
	// 1 - 2 - 4 = -5
	// ( 3 - 2 ** 3 ) * ( -2 ) = 10
	// 3 * ( 3 + 1 * 3 ) / ( -2 ) = -9
	// 3 * ( ( 3 + 1 ) * 3 ) / 2 = 18
	// 1 + 2 ** 10 * 2 = 2049
	// 2 ** 3 mod 3 = 2
	// 2 ** 2 ** 3 = 256
}
