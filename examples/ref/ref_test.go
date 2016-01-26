package ref

import (
	"errors"
	"fmt"
)

func ExampleFieldNames() {
	s := struct {
		id   int
		Name string
		Age  int
	}{}
	fmt.Println(FieldNames(s))
	// Output: [id Name Age] <nil>
}

func ExampleAppendNilError() {
	f := func() {
		fmt.Println("called")
	}
	f2, err := AppendNilError(f, errors.New("test error"))
	fmt.Println("AppendNilError.err:", err)
	fmt.Println(f2.(func() error)())
	// Output:
	// AppendNilError.err: <nil>
	// called
	// test error
}
