package set

import "fmt"

func Example_hasDupeRune() {
	fmt.Println(hasDupeRune("숨바꼭질"))
	fmt.Println(hasDupeRune("다시합시다"))
	// Output:
	// false
	// true
}
