package io

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// bill@mail.com
	// tom@mail.com
	// jane@mail.com
}

func ExampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}
