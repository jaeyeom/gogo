package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("고 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}
