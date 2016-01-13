package set

import "fmt"

func ExampleHasDupeRune() {
	fmt.Println(hasDupeRune("숨바꼭질"))
	fmt.Println(hasDupeRune("다시합시다"))
	// Output:
	// false
	// true
}
