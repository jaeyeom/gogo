package function

import (
	"fmt"
	"strings"
)

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	// Output:
	// [2 3 4 5]
}

func Example_funcLiteral() {
	func() {
		fmt.Println("Hello!")
	}()
	// Output:
	// Hello!
}

func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello!")
	}
	printHello()
	// Output:
	// Hello!
}

func ExampleReadFrom_print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )
}

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}
