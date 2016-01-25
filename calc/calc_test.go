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
	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	fmt.Println(eval("( 3 - 2 ** 3 ) * ( -2 )"))
	fmt.Println(eval("3 * ( 3 + 1 * 3 ) / ( -2 )"))
	fmt.Println(eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 * 2"))
	fmt.Println(eval("2 ** 3 mod 3"))
	fmt.Println(eval("2 ** 2 ** 3"))
	// Output:
	// 5
	// 3
	// -5
	// 10
	// -9
	// 18
	// 2049
	// 2
	// 256
}
